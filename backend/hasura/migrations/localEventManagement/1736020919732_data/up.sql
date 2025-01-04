SET check_function_bodies = false;
CREATE SCHEMA data;
CREATE FUNCTION data.decrease_available_tickets() RETURNS trigger
    LANGUAGE plpgsql
    AS $$
BEGIN
  UPDATE data.events
  SET total_available_tickets = total_available_tickets - 1
  WHERE event_id = NEW.event_id;
  RETURN NEW;
END;
$$;
CREATE FUNCTION data.decrease_followers() RETURNS trigger
    LANGUAGE plpgsql
    AS $$
BEGIN
  UPDATE data.organizations
  SET followers = followers - 1
  WHERE organization_id = OLD.followed_organization_id;
  RETURN OLD;
END;
$$;
CREATE TABLE data.events (
    event_id uuid DEFAULT gen_random_uuid() NOT NULL,
    by_organization_id uuid NOT NULL,
    title character varying NOT NULL,
    ticket_price numeric DEFAULT 0 NOT NULL,
    total_available_tickets integer DEFAULT 0 NOT NULL,
    event_date date,
    description text,
    venue text,
    thumbnail_image_url character varying,
    created_at timestamp without time zone DEFAULT CURRENT_TIMESTAMP,
    category_id uuid,
    last_modified_by uuid,
    CONSTRAINT events_ticket_price_check CHECK ((ticket_price >= (0)::numeric)),
    CONSTRAINT events_total_available_tickets_check CHECK ((total_available_tickets >= 0))
);
CREATE FUNCTION data.get_nearby_events(user_lat double precision, user_lon double precision, radius double precision DEFAULT 100000) RETURNS SETOF data.events
    LANGUAGE plpgsql
    AS $$
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
$$;
CREATE FUNCTION data.increase_followers() RETURNS trigger
    LANGUAGE plpgsql
    AS $$
BEGIN
  UPDATE data.organizations
  SET followers = followers + 1
  WHERE organization_id = NEW.followed_organization_id;
  RETURN NEW;
END;
$$;
CREATE FUNCTION data.insert_into_organizes() RETURNS trigger
    LANGUAGE plpgsql
    AS $$
DECLARE
    id uuid;
BEGIN
    -- Extract the x-hasura-user-id from the session settings as a JSON object and cast it to uuid
    SELECT current_setting('hasura.user')::json ->> 'x-hasura-user-id' INTO id;
    -- Cast the extracted user ID to uuid
    id := id::uuid;
    -- Insert into data.organizes using the fetched user_id and the new organization_id
    INSERT INTO data.organizes (organization_id, organizer_id)
    VALUES (NEW.organization_id, id);
    -- Return the new row
    RETURN NEW;
END;
$$;
CREATE FUNCTION data.track_last_modified_user() RETURNS trigger
    LANGUAGE plpgsql
    AS $$
DECLARE
  user_id uuid;
BEGIN
  SELECT (current_setting('hasura.user')::json ->> 'x-hasura-user-id')::uuid INTO user_id;
  NEW.last_modified_by := user_id;
  RETURN NEW;
END;
$$;
CREATE TABLE data.address (
    addressed_event_id uuid NOT NULL,
    street_name character varying,
    city_name character varying,
    region_name character varying
);
CREATE TABLE data.bookmarks (
    book_marked_event_id uuid NOT NULL,
    book_marker_user_id uuid NOT NULL
);
CREATE TABLE data.categories (
    category_id uuid DEFAULT gen_random_uuid() NOT NULL,
    category_name character varying NOT NULL
);
CREATE TABLE data.event_tags (
    tagged_event_id uuid NOT NULL,
    tag_word_id uuid NOT NULL
);
CREATE TABLE data.follows (
    following_user_id uuid NOT NULL,
    followed_organization_id uuid NOT NULL
);
CREATE TABLE data.images (
    event_id uuid NOT NULL,
    image_url character varying,
    id uuid DEFAULT gen_random_uuid() NOT NULL
);
CREATE TABLE data.locations (
    located_event_id uuid NOT NULL,
    longitude numeric,
    latitude numeric
);
CREATE TABLE data.organizations (
    organization_id uuid DEFAULT gen_random_uuid() NOT NULL,
    organization_name character varying NOT NULL,
    profile_photo_url text,
    bio character varying,
    description text,
    created_at timestamp without time zone DEFAULT CURRENT_TIMESTAMP,
    followers integer DEFAULT 0 NOT NULL,
    CONSTRAINT followers_non_negative CHECK ((followers >= 0))
);
CREATE TABLE data.organizes (
    organizer_id uuid NOT NULL,
    organization_id uuid NOT NULL
);
CREATE TABLE data.payments (
    tx_ref uuid NOT NULL,
    payer_user_id uuid NOT NULL,
    amount real NOT NULL,
    created_at timestamp without time zone DEFAULT CURRENT_TIMESTAMP,
    paid_for_event_id uuid NOT NULL,
    currency character varying(10) DEFAULT 'ETB'::character varying,
    phone_number character varying(15),
    last_name character varying NOT NULL,
    first_name character varying NOT NULL,
    email character varying NOT NULL,
    status character varying,
    chapa_reference character varying,
    chapa_charge real,
    payment_method character varying,
    CONSTRAINT payments_amount_check CHECK ((amount > ((0)::numeric)::double precision))
);
CREATE TABLE data.tags (
    tag_id uuid DEFAULT gen_random_uuid() NOT NULL,
    tag_word character varying NOT NULL
);
CREATE TABLE data.tickets (
    ticket_owner_id uuid NOT NULL,
    tx_ref uuid,
    purchased_date timestamp with time zone DEFAULT now() NOT NULL,
    ticket_id uuid DEFAULT gen_random_uuid() NOT NULL,
    event_id uuid NOT NULL,
    qr_link text
);
CREATE TABLE data.users (
    user_id uuid DEFAULT gen_random_uuid() NOT NULL,
    first_name character varying,
    last_name character varying,
    email character varying NOT NULL,
    password text NOT NULL,
    created_at timestamp without time zone DEFAULT CURRENT_TIMESTAMP,
    role character varying DEFAULT 'user'::character varying NOT NULL
);
ALTER TABLE ONLY data.address
    ADD CONSTRAINT address_pkey PRIMARY KEY (addressed_event_id);
ALTER TABLE ONLY data.bookmarks
    ADD CONSTRAINT bookmarks_pkey PRIMARY KEY (book_marked_event_id, book_marker_user_id);
ALTER TABLE ONLY data.categories
    ADD CONSTRAINT categories_pkey PRIMARY KEY (category_id);
ALTER TABLE ONLY data.event_tags
    ADD CONSTRAINT event_tags_pkey PRIMARY KEY (tagged_event_id, tag_word_id);
ALTER TABLE ONLY data.events
    ADD CONSTRAINT events_pkey PRIMARY KEY (event_id);
ALTER TABLE ONLY data.follows
    ADD CONSTRAINT follows_pkey PRIMARY KEY (following_user_id, followed_organization_id);
ALTER TABLE ONLY data.images
    ADD CONSTRAINT images_pkey PRIMARY KEY (id);
ALTER TABLE ONLY data.locations
    ADD CONSTRAINT locations_pkey PRIMARY KEY (located_event_id);
ALTER TABLE ONLY data.organizations
    ADD CONSTRAINT organizations_pkey PRIMARY KEY (organization_id);
ALTER TABLE ONLY data.organizes
    ADD CONSTRAINT organizes_pkey PRIMARY KEY (organizer_id, organization_id);
ALTER TABLE ONLY data.payments
    ADD CONSTRAINT payments_pkey PRIMARY KEY (tx_ref);
ALTER TABLE ONLY data.tags
    ADD CONSTRAINT tags_pkey PRIMARY KEY (tag_id);
ALTER TABLE ONLY data.tags
    ADD CONSTRAINT tags_tag_word_key UNIQUE (tag_word);
ALTER TABLE ONLY data.tickets
    ADD CONSTRAINT tickets_pkey PRIMARY KEY (ticket_id);
ALTER TABLE ONLY data.tickets
    ADD CONSTRAINT tickets_qr_link_key UNIQUE (qr_link);
ALTER TABLE ONLY data.tickets
    ADD CONSTRAINT tickets_tx_ref_key UNIQUE (tx_ref);
ALTER TABLE ONLY data.categories
    ADD CONSTRAINT unique_category_name UNIQUE (category_name);
ALTER TABLE ONLY data.users
    ADD CONSTRAINT users_email_key UNIQUE (email);
ALTER TABLE ONLY data.users
    ADD CONSTRAINT users_pkey PRIMARY KEY (user_id);
CREATE TRIGGER after_organization_insert AFTER INSERT ON data.organizations FOR EACH ROW EXECUTE FUNCTION data.insert_into_organizes();
CREATE TRIGGER decrease_followers_trigger AFTER DELETE ON data.follows FOR EACH ROW EXECUTE FUNCTION data.decrease_followers();
CREATE TRIGGER decrease_total_available_tickets_trigger AFTER INSERT ON data.tickets FOR EACH ROW EXECUTE FUNCTION data.decrease_available_tickets();
CREATE TRIGGER increase_followers_trigger AFTER INSERT ON data.follows FOR EACH ROW EXECUTE FUNCTION data.increase_followers();
CREATE TRIGGER track_last_modified_user_trigger BEFORE INSERT OR UPDATE ON data.events FOR EACH ROW EXECUTE FUNCTION data.track_last_modified_user();
ALTER TABLE ONLY data.address
    ADD CONSTRAINT address_addressed_event_id_fkey FOREIGN KEY (addressed_event_id) REFERENCES data.events(event_id) ON DELETE CASCADE;
ALTER TABLE ONLY data.bookmarks
    ADD CONSTRAINT bookmarks_book_marked_event_id_fkey FOREIGN KEY (book_marked_event_id) REFERENCES data.events(event_id) ON DELETE CASCADE;
ALTER TABLE ONLY data.bookmarks
    ADD CONSTRAINT bookmarks_book_marker_user_id_fkey FOREIGN KEY (book_marker_user_id) REFERENCES data.users(user_id);
ALTER TABLE ONLY data.event_tags
    ADD CONSTRAINT event_tags_tag_word_id_fkey FOREIGN KEY (tag_word_id) REFERENCES data.tags(tag_id) ON DELETE SET NULL;
ALTER TABLE ONLY data.event_tags
    ADD CONSTRAINT event_tags_tagged_event_id_fkey FOREIGN KEY (tagged_event_id) REFERENCES data.events(event_id) ON DELETE CASCADE;
ALTER TABLE ONLY data.events
    ADD CONSTRAINT events_by_organization_id_fkey FOREIGN KEY (by_organization_id) REFERENCES data.organizations(organization_id) ON DELETE CASCADE;
ALTER TABLE ONLY data.events
    ADD CONSTRAINT events_category_id_fkey FOREIGN KEY (category_id) REFERENCES data.categories(category_id);
ALTER TABLE ONLY data.events
    ADD CONSTRAINT events_user_id_fkey FOREIGN KEY (last_modified_by) REFERENCES data.users(user_id) ON UPDATE CASCADE ON DELETE CASCADE;
ALTER TABLE ONLY data.follows
    ADD CONSTRAINT follows_followed_organization_id_fkey FOREIGN KEY (followed_organization_id) REFERENCES data.organizations(organization_id) ON UPDATE CASCADE ON DELETE CASCADE;
ALTER TABLE ONLY data.follows
    ADD CONSTRAINT follows_following_user_id_fkey FOREIGN KEY (following_user_id) REFERENCES data.users(user_id) ON DELETE CASCADE;
ALTER TABLE ONLY data.images
    ADD CONSTRAINT images_event_id_fkey FOREIGN KEY (event_id) REFERENCES data.events(event_id) ON DELETE CASCADE;
ALTER TABLE ONLY data.locations
    ADD CONSTRAINT locations_located_event_id_fkey FOREIGN KEY (located_event_id) REFERENCES data.events(event_id) ON DELETE CASCADE;
ALTER TABLE ONLY data.organizes
    ADD CONSTRAINT organizes_organization_id_fkey FOREIGN KEY (organization_id) REFERENCES data.organizations(organization_id) ON DELETE CASCADE;
ALTER TABLE ONLY data.organizes
    ADD CONSTRAINT organizes_organizer_id_fkey FOREIGN KEY (organizer_id) REFERENCES data.users(user_id);
ALTER TABLE ONLY data.payments
    ADD CONSTRAINT payments_paid_for_event_id_fkey FOREIGN KEY (paid_for_event_id) REFERENCES data.events(event_id);
ALTER TABLE ONLY data.payments
    ADD CONSTRAINT payments_payer_user_id_fkey FOREIGN KEY (payer_user_id) REFERENCES data.users(user_id);
ALTER TABLE ONLY data.tickets
    ADD CONSTRAINT tickets_event_id_fkey FOREIGN KEY (event_id) REFERENCES data.events(event_id);
ALTER TABLE ONLY data.tickets
    ADD CONSTRAINT tickets_ticket_owner_id_fkey FOREIGN KEY (ticket_owner_id) REFERENCES data.users(user_id);
    
