-- ==============================
-- 🔹 Sample User
-- ==============================

INSERT INTO users (full_name, email, password_hash, phone_number, profile_image, created_at, updated_at)
VALUES (
  'John Doe',
  'john@example.com',
  'hashed_password_here',
  '+94711234567',
  'https://example.com/profile.jpg',
  CURRENT_TIMESTAMP,
  CURRENT_TIMESTAMP
);

-- ==============================
-- 🔹 Sample Properties
-- ==============================

-- Property 1: Sell - House
INSERT INTO properties (
  user_id, title, description, purpose_id, property_type_id, furniture_type_id,
  condition_id, bedrooms, bathrooms, size, size_unit,
  city, address, postal_code, latitude, longitude,
  price, price_unit, is_negotiable, rental_period, is_refundable, pricing_type,
  created_at, updated_at
)
VALUES (
  1, 'Modern 2-Story House', 'A spacious house with garden', 1, 2, 1,
  1, 4, 3, 2500, 'Sqft',
  'Colombo', '123 Lake Road', '00100', 6.9271, 79.8612,
  45000000, 'LKR', TRUE, NULL, FALSE, 'sell',
  CURRENT_TIMESTAMP, CURRENT_TIMESTAMP
);

-- Property 2: Rent - Apartment
INSERT INTO properties (
  user_id, title, description, purpose_id, property_type_id, furniture_type_id,
  condition_id, bedrooms, bathrooms, size, size_unit,
  city, address, postal_code, latitude, longitude,
  price, price_unit, is_negotiable, rental_period, is_refundable, pricing_type,
  created_at, updated_at
)
VALUES (
  1, 'City View Apartment', 'High-rise apartment with balcony', 2, 1, 2,
  2, 2, 1, 950, 'Sqft',
  'Kandy', '45 Temple Street', '20000', 7.2906, 80.6337,
  150000, 'LKR', FALSE, 'Monthly', TRUE, 'rent',
  CURRENT_TIMESTAMP, CURRENT_TIMESTAMP
);

-- Property 3: Stay - Villa
INSERT INTO properties (
  user_id, title, description, purpose_id, property_type_id, furniture_type_id,
  condition_id, bedrooms, bathrooms, size, size_unit,
  city, address, postal_code, latitude, longitude,
  price, price_unit, is_negotiable, rental_period, is_refundable, pricing_type,
  created_at, updated_at
)
VALUES (
  1, 'Beachside Villa', 'Luxury villa near the ocean', 3, 3, 1,
  1, 5, 4, 4000, 'Sqft',
  'Galle', '9 Lighthouse Rd', '80000', 6.0351, 80.2170,
  60000, 'LKR', TRUE, 'Night', FALSE, 'stay',
  CURRENT_TIMESTAMP, CURRENT_TIMESTAMP
);

-- ==============================
-- 🔹 Sample Images
-- ==============================

INSERT INTO property_images (property_id, url, is_primary, created_at, updated_at) VALUES
(1, 'https://example.com/img/house1.jpg', TRUE, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP),
(2, 'https://example.com/img/apartment1.jpg', TRUE, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP),
(3, 'https://example.com/img/villa1.jpg', TRUE, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP);

-- ==============================
-- 🔹 Sample Utilities
-- ==============================

INSERT INTO property_utilities (property_id, utility_id, created_at, updated_at) VALUES
(1, 1, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP), (1, 2, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP), (1, 4, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP),     -- Electricity, Water, Internet
(2, 1, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP), (2, 2, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP),             -- Electricity, Water
(3, 1, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP), (3, 2, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP), (3, 3, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP), (3, 4, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP); -- Electricity, Water, Gas, Internet

-- ==============================
-- 🔹 Sample Amenities
-- ==============================

INSERT INTO property_amenities (property_id, amenity_id, created_at, updated_at) VALUES
(1, 1, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP), (1, 4, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP),             -- Gym, Parking
(2, 1, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP), (2, 6, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP), (2, 7, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP),     -- Gym, Elevator, School Nearby
(3, 2, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP), (3, 4, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP), (3, 5, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP);     -- Swimming Pool, Parking, Security

-- ==============================
-- 🔹 Sample Favourite
-- ==============================

INSERT INTO favourites (user_id, property_id, created_at, updated_at)
VALUES (1, 2, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP);  -- John favorites the apartment
