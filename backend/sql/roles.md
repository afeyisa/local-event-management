# Public User (Unauthenticated)
#       Permissions:
        Browse, filter, and search events by location, date, category, price, title, description, and tags.
        View event details (but not follow, bookmark, or buy tickets).
        Sign up for an account.
        Actions:

        Browse events, but no interaction beyond viewing is allowed.


## Authenticated User (Account Holder)
 #   Permissions:
    Do everything a public user can do.
    
    Follow, bookmark, and buy tickets for events.

    Create, edit, and delete their own events.
    Actions:
    Create an event with the following details:
    Upload multiple images (select featured image).
    Set event location via map.
    Add venue and address.
    Set price (free or specific amount).
    Set event date, category, title, description, and tags.
    Edit and manage only their own events (no access to other users' events).
    Manage tickets they've purchased (view, cancel if allowed).


    to insert into event whith role user
 #       {
 #       "by_organization_id": {
 #       "_eq": "X-Hasura-Organization-Id"
 #      }
 #      }

 # permissions for user on events
 1. select all feilds
 2. update all feild except event id, organization id , created at and event id
 3. delete row
 4. insert all feilds
# permissions for anonymous
1.  select all feilds



user permissions on bookmarks

        {
                "book_marker_user_id": {
                "_eq": "X-Hasura-book_marker_user-Id"
         }
       }