-- ==============================
-- 🔹 purpose_types
-- ==============================

INSERT INTO purpose_types (name, created_at, updated_at) VALUES
('Sell', CURRENT_TIMESTAMP, CURRENT_TIMESTAMP),
('Rent', CURRENT_TIMESTAMP, CURRENT_TIMESTAMP),
('Stay', CURRENT_TIMESTAMP, CURRENT_TIMESTAMP);

-- ==============================
-- 🔹 property_types
-- ==============================

INSERT INTO property_types (name, created_at, updated_at) VALUES
('Apartment', CURRENT_TIMESTAMP, CURRENT_TIMESTAMP),
('House', CURRENT_TIMESTAMP, CURRENT_TIMESTAMP),
('Villa', CURRENT_TIMESTAMP, CURRENT_TIMESTAMP),
('Land', CURRENT_TIMESTAMP, CURRENT_TIMESTAMP),
('Commercial Space', CURRENT_TIMESTAMP, CURRENT_TIMESTAMP),
('Studio', CURRENT_TIMESTAMP, CURRENT_TIMESTAMP);

-- ==============================
-- 🔹 furniture_types
-- ==============================

INSERT INTO furniture_types (name, created_at, updated_at) VALUES
('Fully Furnished', CURRENT_TIMESTAMP, CURRENT_TIMESTAMP),
('Semi Furnished', CURRENT_TIMESTAMP, CURRENT_TIMESTAMP),
('Unfurnished', CURRENT_TIMESTAMP, CURRENT_TIMESTAMP);

-- ==============================
-- 🔹 property_conditions
-- ==============================

INSERT INTO property_conditions (name, created_at, updated_at) VALUES
('New', CURRENT_TIMESTAMP, CURRENT_TIMESTAMP),
('Used', CURRENT_TIMESTAMP, CURRENT_TIMESTAMP),
('Renovated', CURRENT_TIMESTAMP, CURRENT_TIMESTAMP);

-- ==============================
-- 🔹 utilities
-- ==============================

INSERT INTO utilities (name, created_at, updated_at) VALUES
('Electricity', CURRENT_TIMESTAMP, CURRENT_TIMESTAMP),
('Water', CURRENT_TIMESTAMP, CURRENT_TIMESTAMP),
('Gas', CURRENT_TIMESTAMP, CURRENT_TIMESTAMP),
('Internet', CURRENT_TIMESTAMP, CURRENT_TIMESTAMP),
('Sewage', CURRENT_TIMESTAMP, CURRENT_TIMESTAMP),
('Air Conditioning', CURRENT_TIMESTAMP, CURRENT_TIMESTAMP);

-- ==============================
-- 🔹 amenities
-- ==============================

INSERT INTO amenities (name, created_at, updated_at) VALUES
('Gym', CURRENT_TIMESTAMP, CURRENT_TIMESTAMP),
('Swimming Pool', CURRENT_TIMESTAMP, CURRENT_TIMESTAMP),
('Playground', CURRENT_TIMESTAMP, CURRENT_TIMESTAMP),
('Parking', CURRENT_TIMESTAMP, CURRENT_TIMESTAMP),
('Security', CURRENT_TIMESTAMP, CURRENT_TIMESTAMP),
('Elevator', CURRENT_TIMESTAMP, CURRENT_TIMESTAMP),
('School Nearby', CURRENT_TIMESTAMP, CURRENT_TIMESTAMP),
('Hospital Nearby', CURRENT_TIMESTAMP, CURRENT_TIMESTAMP),
('Public Transport Access', CURRENT_TIMESTAMP, CURRENT_TIMESTAMP);
