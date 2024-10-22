CREATE TABLE data.users (
  user_id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
  first_name varchar,
  last_name varchar,
  email varchar NOT NULL UNIQUE,
  password text NOT NULL,
  created_at timestamp DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE data.organizes (
  organizer_id uuid NOT NULL,
  organization_id uuid NOT NULL,
  PRIMARY KEY (organizer_id, organization_id)
);

CREATE TABLE data.organizations (
  organization_id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
  organization_name varchar NOT NULL,
  profile_photo_url text,
  bio varchar,
  description text,
  created_at timestamp DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE data.follows (
  following_user_id uuid NOT NULL,
  followed_organization_id uuid NOT NULL,
  PRIMARY KEY (following_user_id, followed_organization_id)
);

CREATE TABLE data.events (
  event_id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
  by_organization_id uuid NOT NULL,
  title varchar NOT NULL,
  ticket_price decimal DEFAULT 0 CHECK (ticket_price >= 0),
  total_available_tickets int DEFAULT 0 CHECK (total_available_tickets >= 0),
  event_date date,
  category varchar,
  description text,
  is_online boolean,
  venue text,
  thumbnail_image_url varchar,
  created_at timestamp DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE categories (
  category_id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
  category_name varchar NOT NULL
);


CREATE TABLE data.images (
  event_id uuid NOT NULL,
  image_url varchar NOT NULL,
  PRIMARY KEY (event_id, image_url)
);

CREATE TABLE data.locations (
  located_event_id uuid PRIMARY KEY NOT NULL,
  longitude decimal,
  latitude decimal
);

CREATE TABLE data.address (
  addressed_event_id uuid PRIMARY KEY NOT NULL,
  street_name varchar,
  city_name varchar,
  region_name varchar
);

CREATE TABLE data.tags (
  tag_id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
  tag_word varchar UNIQUE NOT NULL
);

CREATE TABLE data.event_tags (
  tagged_event_id uuid NOT NULL,
  tag_word_id uuid NOT NULL,
  PRIMARY KEY (tagged_event_id, tag_word_id)
);

CREATE TABLE data.bookmarks (
  book_marked_event_id uuid NOT NULL,
  book_marker_user_id uuid NOT NULL,
  PRIMARY KEY (book_marked_event_id, book_marker_user_id)
);

CREATE TABLE data.payments (
  transaction_no varchar PRIMARY KEY NOT NULL,
  payer_user_id uuid NOT NULL,
  amount decimal CHECK (amount > 0),
  transaction_time timestamp DEFAULT CURRENT_TIMESTAMP,
  paid_for_event_id uuid NOT NULL
);

CREATE TABLE data.tickets (
  ticket_id uuid PRIMARY KEY DEFAULT gen_random_uuid(),
  for_event_id uuid NOT NULL,
  transaction_no varchar NOT NULL
);

CREATE TABLE data.notifications (
  notification_id uuid PRIMARY KEY DEFAULT gen_random_uuid(),
  user_id uuid NOT NULL,
  message text,
  is_read boolean,
  created_at timestamp DEFAULT CURRENT_TIMESTAMP
);



ALTER TABLE data.categories
  ADD CONSTRAINT unique_category_name UNIQUE (category_name);

ALTER TABLE data.organizes 
ADD FOREIGN KEY (organizer_id) REFERENCES data.users (user_id);
ALTER TABLE data.events ADD COLUMN category_id uuid REFERENCES categories(category_id);
ALTER TABLE data.events ADD COLUMN is_online boolean  DEFAULT FALSE;
ALTER TABLE data.organizes 
ADD FOREIGN KEY (organization_id) REFERENCES data.organizations (organization_id);

ALTER TABLE data.follows 
ADD FOREIGN KEY (following_user_id) REFERENCES data.users (user_id);

ALTER TABLE data.follows 
ADD FOREIGN KEY (followed_organization_id) REFERENCES data.organizations (organization_id);

ALTER TABLE data.events 
ADD FOREIGN KEY (by_organization_id) REFERENCES data.organizations (organization_id);

ALTER TABLE data.images 
ADD FOREIGN KEY (event_id) REFERENCES data.events (event_id);

ALTER TABLE data.locations 
ADD FOREIGN KEY (located_event_id) REFERENCES data.events (event_id);

ALTER TABLE data.address 
ADD FOREIGN KEY (addressed_event_id) REFERENCES data.events (event_id);

ALTER TABLE data.event_tags 
ADD FOREIGN KEY (tagged_event_id) REFERENCES data.events (event_id);

ALTER TABLE data.event_tags 
ADD FOREIGN KEY (tag_word_id) REFERENCES data.tags (tag_id);

ALTER TABLE data.bookmarks 
ADD FOREIGN KEY (book_marked_event_id) REFERENCES data.events (event_id);

ALTER TABLE data.bookmarks 
ADD FOREIGN KEY (book_marker_user_id) REFERENCES data.users (user_id);

ALTER TABLE data.payments 
ADD FOREIGN KEY (payer_user_id) REFERENCES data.users (user_id);

ALTER TABLE data.payments 
ADD FOREIGN KEY (paid_for_event_id) REFERENCES data.events (event_id);

ALTER TABLE data.tickets 
ADD FOREIGN KEY (for_event_id) REFERENCES data.events (event_id);

ALTER TABLE data.tickets 
ADD FOREIGN KEY (ticket_owner_id) REFERENCES data.users (user_id);

ALTER TABLE data.tickets 
ADD FOREIGN KEY (transaction_no) REFERENCES data.payments (transaction_no);

ALTER TABLE data.notifications 
ADD FOREIGN KEY (user_id) REFERENCES data.users (user_id);


-- Insert categories
INSERT INTO categories (category_name) VALUES
  ( 'Technology'),
  ( 'Music'),
  ( 'Education'),
  ( 'Sports'),
  ( 'Art'),
  ( 'Health'),
  ( 'Business'),
  ( 'Entertainment'),
  ( 'Science'),
  ( 'Travel');

-- populating events if catagory id don't known
  INSERT INTO events (
  event_id, 
  by_organization_id, 
  title, 
  ticket_price, 
  total_available_tickets, 
  event_date, 
  category_id, 
  description, 
  venue, 
  is_online, 
  thumnail_image_url, 
  created_at
) 
VALUES (
  uuid_generate_v4(),  -- Generates a new unique event_id
  'c5b8efdc-5f65-4ef8-95e5-7f4edb789c24',  -- Organization ID
  'Music Festival 2024',  -- Event title
  75.00,  -- Ticket price
  1000,  -- Total available tickets
  '2024-11-05',  -- Event date
  (SELECT category_id FROM categories WHERE category_name = 'Music'),  -- Subquery to get category_id
  'A grand music festival with live performances from top artists.',  -- Description
  'Central Park, New York',  -- Venue
  false,  -- is_online (false = in-person)
  'https://example.com/music-festival-thumbnail.jpg',  -- Thumbnail image URL
  NOW()  -- Created at (current timestamp)
);


CREATE OR REPLACE FUNCTION data.get_user_organizations(user_id uuid)
RETURNS SETOF uuid AS $$
  SELECT organization_id
  FROM data.organizes
  WHERE organizer_id = user_id;
$$ LANGUAGE sql STABLE;
DROP FUNCTION IF EXISTS data.get_user_organizations(uuid);



query BrowseEvents(
  $location: String, 
  $start_date: date, 
  $end_date: date, 
  $category_id: uuid, 
  $min_price: numeric, 
  $max_price: numeric
) {
  events(
    where: {
      _and: [
        { is_public: { _eq: true } },  # Only show public events
        { venue: { _ilike: $location } },  # Filter by location (partial match)
        { event_date: { _gte: $start_date } },  # Filter events starting from a specific date
        { event_date: { _lte: $end_date } },  # Filter events until a specific date
        { category_id: { _eq: $category_id } },  # Filter by category
        { ticket_price: { _gte: $min_price } },  # Filter events with price >= min_price
        { ticket_price: { _lte: $max_price } }  # Filter events with price <= max_price
      ]
    },
    order_by: { event_date: asc },  # Sort events by date
  ) {
    event_id
    title
    event_date
    ticket_price
    description
    venue
    category {
      category_name
    }
    thumbnail_image_url
  }
}

CREATE TYPE data.organize_result AS (
    organization_id uuid,
    organizer_id uuid
);

DROP FUNCTION IF EXISTS data.insert_into_organizes();
CREATE OR REPLACE FUNCTION data.insert_into_organizes()
RETURNS TABLE(organization_id uuid, organizer_id uuid)
LANGUAGE plpgsql
AS $function$
DECLARE
    user_id uuid;
BEGIN
    -- Fetch the user_id (organizer_id) dynamically from the users table or JWT session
    SELECT user_id INTO user_id
    FROM data_users
    WHERE user_id = current_setting('hasura.user')::uuid;

    -- Insert into data_organizes using the fetched user_id and the new organization_id
    INSERT INTO data.data_organizes (organization_id, organizer_id)
    VALUES (NEW.organization_id, user_id);

    -- Return the inserted data as a table
    RETURN QUERY SELECT NEW.organization_id, user_id;
END;
$function$;



CREATE TRIGGER after_organization_insert
AFTER INSERT ON data.organizations
FOR EACH ROW
EXECUTE FUNCTION data.insert_into_organizes();


DROP TRIGGER after_organization_insert ON data.organizations;
DROP FUNCTION IF EXISTS data.insert_into_organizes();


DROP FUNCTION IF EXISTS data.insert_into_organizes();
CREATE OR REPLACE FUNCTION data.insert_into_organizes()
RETURNS TRIGGER
LANGUAGE plpgsql
AS $function$
DECLARE
    user_id uuid;
BEGIN
    -- Fetch the user_id (organizer_id) dynamically from the users table or JWT session
    SELECT user_id INTO user_id
    FROM data.data_users
    WHERE user_id = current_setting('hasura.user');

    -- Insert into data_organizes using the fetched user_id and the new organization_id
    INSERT INTO data.data_organizes (organization_id, organizer_id)
    VALUES (NEW.organization_id, user_id);

    -- Return the new row (or NULL if not applicable)
    RETURN NEW;
END;
$function$;

CREATE TRIGGER after_organization_insert
AFTER INSERT ON data.organizations
FOR EACH ROW
EXECUTE FUNCTION data.insert_into_organizes();

