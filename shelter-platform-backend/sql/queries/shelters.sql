-- name: CreateShelter :one
   INSERT INTO shelters (
       name,
       city,
       state,
       country,
       zip_code,
       phone_number,
       contact_email,
       founded_at
   ) VALUES (
       $1,
       $2,
       $3,
       $4,
       $5,
       $6,
       $7,
       $8
   )
   RETURNING
       id,
       name,
       city,
       state,
       country,
       zip_code,
       phone_number,
       contact_email,
       founded_at,
       created_at,
       updated_at;