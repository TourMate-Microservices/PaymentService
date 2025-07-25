-- DROP TABLES IF EXIST (optional for re-run)
DROP TABLE IF EXISTS feedbacks;
DROP TABLE IF EXISTS payments;

-- ===============================
-- ✅ Create feedbacks table
-- ===============================
CREATE TABLE feedbacks (
    feedback_id SERIAL PRIMARY KEY,
    customer_id INT NOT NULL,
    service_id INT NOT NULL,
    tour_guide_id INT NOT NULL,
    created_date TIMESTAMP NOT NULL,
    content TEXT NOT NULL,
    rating INT NOT NULL CHECK (rating >= 1 AND rating <= 5),
    is_deleted BOOLEAN NOT NULL DEFAULT FALSE,
    updated_at TIMESTAMP NOT NULL,
    invoice_id INT NOT NULL
);

-- ✅ Seed data for feedbacks
INSERT INTO feedbacks (customer_id, service_id, tour_guide_id, created_date, content, rating, is_deleted, updated_at, invoice_id) VALUES
(1, 101, 501, NOW() - INTERVAL '10 days', 'Amazing trip! Highly recommend.', 5, FALSE, NOW(), 1001),
(2, 102, 502, NOW() - INTERVAL '8 days', 'Good experience, but weather was bad.', 4, FALSE, NOW(), 1002),
(3, 103, 503, NOW() - INTERVAL '6 days', 'Not satisfied with the guide.', 2, FALSE, NOW(), 1003),
(4, 101, 501, NOW() - INTERVAL '4 days', 'Nice guide and smooth schedule.', 4, FALSE, NOW(), 1004),
(5, 104, 504, NOW() - INTERVAL '2 days', 'Excellent service overall!', 5, FALSE, NOW(), 1005);

-- ===============================
-- ✅ Create payments table
-- ===============================
CREATE TABLE payments (
    payment_id SERIAL PRIMARY KEY,
    price REAL NOT NULL CHECK (price >= 0),
    status VARCHAR(50) NOT NULL,
    created_at TIMESTAMP NOT NULL,
    payment_method VARCHAR(50) NOT NULL,
    customer_id INT NOT NULL,
    invoice_id INT
);

-- ✅ Seed data for payments
INSERT INTO payments (price, status, created_at, payment_method, customer_id, invoice_id) VALUES
(150.0, 'Completed', NOW() - INTERVAL '10 days', 'Credit Card', 1, 1001),
(200.5, 'Pending', NOW() - INTERVAL '9 days', 'PayPal', 2, 1002),
(120.0, 'Completed', NOW() - INTERVAL '8 days', 'Cash', 3, 1003),
(175.75, 'Failed', NOW() - INTERVAL '7 days', 'Bank Transfer', 4, 1004),
(300.0, 'Completed', NOW() - INTERVAL '6 days', 'Credit Card', 5, 1005);
