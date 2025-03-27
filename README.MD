### 1. **Users Table**
This table stores user information such as name, phone number, and status.

```sql
CREATE TABLE users (
    user_id INT AUTO_INCREMENT PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    phone_number VARCHAR(15) UNIQUE NOT NULL,
    status VARCHAR(255),
    profile_picture_url VARCHAR(255),
    last_seen TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);
```

**Example Data:**

```sql
INSERT INTO users (name, phone_number, status, profile_picture_url) VALUES
('John Doe', '1234567890', 'Hey there! I am using WhatsApp.', 'https://example.com/profile1.jpg'),
('Jane Smith', '0987654321', 'Busy', 'https://example.com/profile2.jpg');
```

### 2. **Chats Table**
This table stores information about a chat. A chat could be a group or a one-on-one chat.

```sql
CREATE TABLE chats (
    chat_id INT AUTO_INCREMENT PRIMARY KEY,
    chat_name VARCHAR(255),
    is_group BOOLEAN DEFAULT FALSE,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);
```

**Example Data:**

```sql
INSERT INTO chats (chat_name, is_group) VALUES
('Family Group', TRUE),
('John and Jane', FALSE);
```

### 3. **Chat Participants Table**
This table keeps track of users that are part of a chat.

```sql
CREATE TABLE chat_participants (
    chat_id INT,
    user_id INT,
    join_date TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY (chat_id, user_id),
    FOREIGN KEY (chat_id) REFERENCES chats(chat_id),
    FOREIGN KEY (user_id) REFERENCES users(user_id)
);
```

**Example Data:**

```sql
INSERT INTO chat_participants (chat_id, user_id) VALUES
(1, 1), -- John is part of Family Group
(1, 2), -- Jane is part of Family Group
(2, 1), -- John is part of John and Jane chat
(2, 2); -- Jane is part of John and Jane chat
```

### 4. **Messages Table**
This table stores the messages sent by users in each chat.

```sql
CREATE TABLE messages (
    message_id INT AUTO_INCREMENT PRIMARY KEY,
    chat_id INT,
    sender_id INT,
    message_text TEXT,
    message_type ENUM('text', 'image', 'video', 'audio', 'document') DEFAULT 'text',
    media_url VARCHAR(255),
    sent_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (chat_id) REFERENCES chats(chat_id),
    FOREIGN KEY (sender_id) REFERENCES users(user_id)
);
```

**Example Data:**

```sql
INSERT INTO messages (chat_id, sender_id, message_text, message_type, media_url) VALUES
(1, 1, 'Hello family!', 'text', NULL),
(1, 2, 'Hey John!', 'text', NULL),
(2, 1, 'Hi Jane, how are you?', 'text', NULL);
```

### 5. **Media Table**
This table stores information about media files (images, videos, etc.) uploaded by users in the chats.

```sql
CREATE TABLE media (
    media_id INT AUTO_INCREMENT PRIMARY KEY,
    message_id INT,
    media_type ENUM('image', 'video', 'audio', 'document') NOT NULL,
    media_url VARCHAR(255) NOT NULL,
    uploaded_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (message_id) REFERENCES messages(message_id)
);
```

**Example Data:**

```sql
INSERT INTO media (message_id, media_type, media_url) VALUES
(1, 'image', 'https://example.com/image1.jpg'),
(2, 'video', 'https://example.com/video1.mp4');
```

### 6. **Message Status Table**
This table keeps track of message statuses (sent, delivered, read).

```sql
CREATE TABLE message_status (
    message_id INT,
    user_id INT,
    status ENUM('sent', 'delivered', 'read') NOT NULL,
    status_updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY (message_id, user_id),
    FOREIGN KEY (message_id) REFERENCES messages(message_id),
    FOREIGN KEY (user_id) REFERENCES users(user_id)
);
```

**Example Data:**

```sql
INSERT INTO message_status (message_id, user_id, status) VALUES
(1, 1, 'sent'),
(1, 2, 'delivered'),
(2, 2, 'read');
```

### 7. **Contacts Table**
This table stores the contact information between users, allowing them to see their friends.

```sql
CREATE TABLE contacts (
    user_id INT,
    contact_id INT,
    added_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY (user_id, contact_id),
    FOREIGN KEY (user_id) REFERENCES users(user_id),
    FOREIGN KEY (contact_id) REFERENCES users(user_id)
);
```

**Example Data:**

```sql
INSERT INTO contacts (user_id, contact_id) VALUES
(1, 2), -- John has Jane as a contact
(2, 1); -- Jane has John as a contact
```

### 8. **Blocking Table (optional)**
A table to manage users who block others.

```sql
CREATE TABLE blocks (
    blocker_id INT,
    blocked_id INT,
    blocked_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY (blocker_id, blocked_id),
    FOREIGN KEY (blocker_id) REFERENCES users(user_id),
    FOREIGN KEY (blocked_id) REFERENCES users(user_id)
);
```

**Example Data:**

```sql
INSERT INTO blocks (blocker_id, blocked_id) VALUES
(1, 2); -- John has blocked Jane
```

---
