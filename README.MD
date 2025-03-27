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



### Scenario Overview:
1. **User clicks on a user profile**: You will create a chat between the two users if one doesn't exist already.
2. **User clicks on "Chat"**: A new entry in the `chats` table is created for the two users (this would be a one-on-one chat, as opposed to a group chat).
3. **Send a message**: When a message is sent, the following things happen:
   - **`chat_participants`**: The participants (users) are associated with the chat in the `chat_participants` table (this can happen when you first create the chat, or if they are not yet part of the chat).
   - **`messages`**: A new record is created in the `messages` table with the message content, sender, and chat association.

### Detailed Workflow:

1. **Clicking on User Profile & Creating a Chat**:
   - When User A clicks on User B's profile to start a chat, you check if a chat already exists between these two users.
   - If **no chat exists**, you create a new record in the `chats` table (for a one-on-one chat), and associate both users as participants in that chat in the `chat_participants` table.

   **SQL Workflow:**

   - **Create a new chat** (if none exists):
     ```sql
     INSERT INTO chats (chat_name, is_group) VALUES ('User A and User B', FALSE);
     ```

   - **Create participants** (add both users to the chat):
     ```sql
     INSERT INTO chat_participants (chat_id, user_id) VALUES
     (new_chat_id, user_a_id),
     (new_chat_id, user_b_id);
     ```

2. **Sending a Message**:
   - Once the chat exists, when one of the users (say User A) sends a message, a record is created in the `messages` table.
   - The message is associated with the chat (via `chat_id`), and the sender is recorded (via `sender_id`).

   **SQL Workflow:**

   - **Insert a message**:
     ```sql
     INSERT INTO messages (chat_id, sender_id, message_text, message_type) VALUES
     (existing_chat_id, user_a_id, 'Hello!', 'text');
     ```

   - If there are media (e.g., an image), you would also insert the media data into the `media` table (if you have this feature).

   - **Track message status**:
     When the message is sent, you might want to track its status (sent, delivered, read, etc.) in the `message_status` table.
     ```sql
     INSERT INTO message_status (message_id, user_id, status) VALUES
     (message_id, user_b_id, 'delivered'); -- For example, when User B receives the message
     ```

### Diagram of the Flow:

1. **User clicks on another user** -> **Check if chat exists**:
   - If no chat exists, create a new chat.
   - Insert records into `chats` and `chat_participants` for both users.

2. **User sends a message** -> **Insert the message into `messages` table**:
   - Store the message content and the sender's ID.
   - Optionally, if media is sent, store media info in the `media` table.
   - Track the message's delivery status in the `message_status` table.

---

### Example SQL Workflow (step-by-step):
Let's assume **User A** sends a message to **User B**.

1. **Check if chat exists**:

```sql
SELECT chat_id FROM chats
WHERE (chat_id IN (SELECT chat_id FROM chat_participants WHERE user_id = 1)
        AND chat_id IN (SELECT chat_id FROM chat_participants WHERE user_id = 2));
```

If **no chat** exists, create a new chat:
```sql
INSERT INTO chats (chat_name, is_group) VALUES ('User A and User B', FALSE);
```

Assume the `chat_id` for this new chat is `1001`.

2. **Add participants to the new chat**:

```sql
INSERT INTO chat_participants (chat_id, user_id) VALUES
(1001, 1), -- User A
(1001, 2); -- User B
```

3. **Send a message**:
User A sends a message:
```sql
INSERT INTO messages (chat_id, sender_id, message_text, message_type) VALUES
(1001, 1, 'Hi, User B!', 'text');
```

4. **Message Status**:
Track the message delivery status:
```sql
-- Assuming message_id is 10001 (auto-generated after insertion):
INSERT INTO message_status (message_id, user_id, status) VALUES
(10001, 2, 'delivered'); -- When User B receives the message
```

---

### Summary:

- **Step 1**: When you click on a user's profile to chat, you check if a chat exists between them. If not, you create a new `chats` entry and add both users as participants in `chat_participants`.
- **Step 2**: When a user sends a message, a new message is inserted into the `messages` table, and the status of the message can be tracked in the `message_status` table.

This workflow ensures that a new chat is created only once and that messages are properly tracked within the system.