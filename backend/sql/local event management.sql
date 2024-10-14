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
  venue text,
  event_type varchar,
  thumbnail_image_url varchar,
  created_at timestamp DEFAULT CURRENT_TIMESTAMP
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



ALTER TABLE data.organizes 
ADD FOREIGN KEY (organizer_id) REFERENCES data.users (user_id);

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
ADD FOREIGN KEY (transaction_no) REFERENCES data.payments (transaction_no);

ALTER TABLE data.notifications 
ADD FOREIGN KEY (user_id) REFERENCES data.users (user_id);
