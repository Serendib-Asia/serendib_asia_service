-- ==============================
-- 🔹 Sample User
-- ==============================

INSERT INTO users (full_name, email, password_hash, phone_number, profile_image)
VALUES (
  'John Doe',
  'john@example.com',
  'hashed_password_here',
  '+94711234567',
  'https://example.com/profile.jpg'
);

-- ==============================
-- 🔹 Sample Properties
-- ==============================

-- Property 1: Sell - House
INSERT INTO properties (
  user_id, title, description, purpose_id, property_type_id, furniture_type_id,
  condition_id, bedrooms, bathrooms, size, size_unit,
  city, address, postal_code, latitude, longitude,
  price, price_unit, is_negotiable, rental_period, is_refundable, pricing_type
)
VALUES (
  1, 'Modern 2-Story House', 'A spacious house with garden', 1, 2, 1,
  1, 4, 3, 2500, 'Sqft',
  'Colombo', '123 Lake Road', '00100', 6.9271, 79.8612,
  45000000, 'LKR', TRUE, NULL, FALSE, 'sell'
);

-- Property 2: Rent - Apartment
INSERT INTO properties (
  user_id, title, description, purpose_id, property_type_id, furniture_type_id,
  condition_id, bedrooms, bathrooms, size, size_unit,
  city, address, postal_code, latitude, longitude,
  price, price_unit, is_negotiable, rental_period, is_refundable, pricing_type
)
VALUES (
  1, 'City View Apartment', 'High-rise apartment with balcony', 2, 1, 2,
  2, 2, 1, 950, 'Sqft',
  'Kandy', '45 Temple Street', '20000', 7.2906, 80.6337,
  150000, 'LKR', FALSE, 'Monthly', TRUE, 'rent'
);

-- Property 3: Stay - Villa
INSERT INTO properties (
  user_id, title, description, purpose_id, property_type_id, furniture_type_id,
  condition_id, bedrooms, bathrooms, size, size_unit,
  city, address, postal_code, latitude, longitude,
  price, price_unit, is_negotiable, rental_period, is_refundable, pricing_type
)
VALUES (
  1, 'Beachside Villa', 'Luxury villa near the ocean', 3, 3, 1,
  1, 5, 4, 4000, 'Sqft',
  'Galle', '9 Lighthouse Rd', '80000', 6.0351, 80.2170,
  60000, 'LKR', TRUE, 'Night', FALSE, 'stay'
);

-- ==============================
-- 🔹 Sample Images
-- ==============================

INSERT INTO property_images (property_id, url, is_primary) VALUES
(1, 'https://example.com/img/house1.jpg', TRUE),
(2, 'https://example.com/img/apartment1.jpg', TRUE),
(3, 'https://example.com/img/villa1.jpg', TRUE);

-- ==============================
-- 🔹 Sample Utilities
-- ==============================

INSERT INTO property_utilities (property_id, utility_id) VALUES
(1, 1), (1, 2), (1, 4),     -- Electricity, Water, Internet
(2, 1), (2, 2),             -- Electricity, Water
(3, 1), (3, 2), (3, 3), (3, 4); -- Electricity, Water, Gas, Internet

-- ==============================
-- 🔹 Sample Amenities
-- ==============================

INSERT INTO property_amenities (property_id, amenity_id) VALUES
(1, 1), (1, 4),             -- Gym, Parking
(2, 1), (2, 6), (2, 7),     -- Gym, Elevator, School Nearby
(3, 2), (3, 4), (3, 5);     -- Swimming Pool, Parking, Security

-- ==============================
-- 🔹 Sample Favourite
-- ==============================

INSERT INTO favourites (user_id, property_id)
VALUES (1, 2);  -- John favorites the apartment
