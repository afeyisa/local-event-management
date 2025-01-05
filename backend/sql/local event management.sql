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
  followers int DEFAULT 0
  created_at timestamp DEFAULT CURRENT_TIMESTAMP
  CHECK(followers >= 0)
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
  category_id varchar,
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
  id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
  event_id uuid NOT NULL,
  image_url varchar ,
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
  book_marked_event_id uuid unique NOT NULL,
  book_marker_user_id uuid unique NOT NULL,
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

CREATE OR REPLACE FUNCTION data.increase_followers()
RETURNS TRIGGER AS $$
BEGIN
  UPDATE data.organizations
  SET followers = followers + 1
  WHERE organization_id = NEW.followed_organization_id;
  RETURN NEW;
END;
$$ LANGUAGE plpgsql;

CREATE OR REPLACE FUNCTION data.decrease_followers()
RETURNS TRIGGER AS $$
BEGIN
  UPDATE data.organizations
  SET followers = followers - 1
  WHERE organization_id = OLD.followed_organization_id;
  RETURN OLD;
END;
$$ LANGUAGE plpgsql;

CREATE TRIGGER increase_followers_trigger
AFTER INSERT ON data.follows
FOR EACH ROW
EXECUTE FUNCTION data.increase_followers();
-- decrease number avaliable tickets
CREATE OR REPLACE FUNCTION data.decrease_available_tickets()
RETURNS TRIGGER AS $$
BEGIN
  UPDATE data.events
  SET total_available_tickets = total_available_tickets - 1
  WHERE event_id = NEW.event_id;
  RETURN NEW;
END;
$$ LANGUAGE plpgsql;

CREATE TRIGGER decrease_total_available_tickets_trigger
AFTER INSERT ON data.tickets
FOR EACH ROW
EXECUTE FUNCTION data.decrease_available_tickets();
-- -------------------------------------------
CREATE TRIGGER decrease_followers_trigger
AFTER DELETE ON data.follows
FOR EACH ROW
EXECUTE FUNCTION data.decrease_followers();


-- CREATE OR REPLACE FUNCTION data.get_user_organizations(user_id uuid)
-- RETURNS SETOF uuid AS $$
--   SELECT organization_id
--   FROM data.organizes
--   WHERE organizer_id = user_id;
-- $$ LANGUAGE sql STABLE;
-- DROP FUNCTION IF EXISTS data.get_user_organizations(uuid);

-- DROP FUNCTION IF EXISTS data.insert_into_organizes();

CREATE OR REPLACE FUNCTION data.insert_into_organizes()
RETURNS TRIGGER
LANGUAGE plpgsql
AS $function$
DECLARE
    user_id uuid;
BEGIN
    SELECT user_id INTO user_id
    FROM data.data_users
    WHERE user_id = current_setting('hasura.user');
    INSERT INTO data.data_organizes (organization_id, organizer_id)
    VALUES (NEW.organization_id, user_id);
    RETURN NEW;
END;
$function$;

CREATE TRIGGER after_organization_insert
AFTER INSERT ON data.organizations
FOR EACH ROW
EXECUTE FUNCTION data.insert_into_organizes();

CREATE EXTENSION IF NOT EXISTS postgis SCHEMA data;
CREATE OR REPLACE FUNCTION data.get_nearby_events(user_lat double precision, user_lon double precision, radius double precision DEFAULT 100000)
 RETURNS SETOF data.events
 LANGUAGE plpgsql
AS $function$
BEGIN
  RETURN QUERY
  SELECT e.*
  FROM data.events e
  JOIN data.locations l ON e.event_id = l.located_event_id
  WHERE data.ST_DWithin(
         data.ST_SetSRID(data.ST_MakePoint(l.longitude, l.latitude), 4326)::data.geography, 
         data.ST_SetSRID(data.ST_MakePoint(user_lon, user_lat), 4326)::data.geography,
         radius
      )
  ORDER BY data.ST_Distance(
            data.ST_SetSRID(data.ST_MakePoint(l.longitude, l.latitude), 4326)::data.geography, 
            data.ST_SetSRID(data.ST_MakePoint(user_lon, user_lat), 4326)::data.geography
         ) ASC;  
END;
$function$

CREATE OR REPLACE FUNCTION track_last_modified_user()
RETURNS TRIGGER AS $$
DECLARE
  user_id uuid;
BEGIN
  SELECT (current_setting('hasura.user')::json ->> 'x-hasura-user-id')::uuid INTO user_id;
  NEW.last_modified_by := user_id;
  RETURN NEW;
END;
$$ LANGUAGE plpgsql;


CREATE TRIGGER track_last_modified_user_trigger
BEFORE INSERT OR UPDATE ON events
FOR EACH ROW
EXECUTE FUNCTION data.track_last_modified_user();
