
### 1. **Users Table**
This table holds the information for all the users.

```sql
CREATE TABLE Users (
    user_id INT PRIMARY KEY AUTO_INCREMENT,  -- Unique identifier for each user
    name VARCHAR(100) NOT NULL,              -- User's name
    phone_number VARCHAR(15) NOT NULL UNIQUE, -- User's phone number (unique)
    profile_picture VARCHAR(255),            -- URL to user's profile picture
    status VARCHAR(255),                     -- User's status (optional)
    last_seen TIMESTAMP,                     -- Last seen timestamp
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP  -- Account creation timestamp
);
```

### 2. **Chats Table**
This table stores the chat information between users.

```sql
CREATE TABLE Chats (
    chat_id INT PRIMARY KEY AUTO_INCREMENT,   -- Unique identifier for each chat
    chat_name VARCHAR(255),                   -- Name of the chat (for group chats)
    is_group BOOLEAN DEFAULT FALSE,           -- Whether the chat is a group chat
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP -- Time the chat was created
);
```

### 3. **UserChats Table**
This table keeps track of which users are part of which chat (many-to-many relationship).

```sql
CREATE TABLE UserChats (
    user_id INT,
    chat_id INT,
    is_admin BOOLEAN DEFAULT FALSE,         -- Whether the user is an admin of the group
    joined_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,  -- Time the user joined the chat
    PRIMARY KEY (user_id, chat_id),
    FOREIGN KEY (user_id) REFERENCES Users(user_id) ON DELETE CASCADE,
    FOREIGN KEY (chat_id) REFERENCES Chats(chat_id) ON DELETE CASCADE
);
```

### 4. **Messages Table**
This table stores all the messages exchanged within the chats.

```sql
CREATE TABLE Messages (
    message_id INT PRIMARY KEY AUTO_INCREMENT, -- Unique identifier for each message
    chat_id INT,                               -- The chat where the message was sent
    sender_id INT,                             -- The user who sent the message
    message_text TEXT,                         -- The content of the message
    media_url VARCHAR(255),                    -- URL to media (if the message contains media)
    message_type ENUM('text', 'image', 'video', 'document') NOT NULL,  -- Message type
    timestamp TIMESTAMP DEFAULT CURRENT_TIMESTAMP, -- When the message was sent
    is_read BOOLEAN DEFAULT FALSE,             -- Whether the message has been read by the recipient
    FOREIGN KEY (chat_id) REFERENCES Chats(chat_id) ON DELETE CASCADE,
    FOREIGN KEY (sender_id) REFERENCES Users(user_id) ON DELETE CASCADE
);
```

### 5. **Media Table**
This table stores media files sent in the messages, like images, videos, or documents.

```sql
CREATE TABLE Media (
    media_id INT PRIMARY KEY AUTO_INCREMENT,   -- Unique identifier for the media
    message_id INT,                            -- The message this media is attached to
    media_url VARCHAR(255),                    -- URL to the media
    media_type ENUM('image', 'video', 'document') NOT NULL, -- Type of media
    uploaded_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP, -- When the media was uploaded
    FOREIGN KEY (message_id) REFERENCES Messages(message_id) ON DELETE CASCADE
);
```

### 6. **Groups Table** (Optional, if you'd like to separate groups into a table)
This table stores information specific to groups (if using separate from `Chats` table).

```sql
CREATE TABLE Groups (
    group_id INT PRIMARY KEY AUTO_INCREMENT,  -- Unique identifier for each group
    group_name VARCHAR(255),                   -- Group name
    group_icon VARCHAR(255),                   -- URL to group icon
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP -- When the group was created
);
```

### 7. **GroupMembers Table**
This table holds information about members of the groups.

```sql
CREATE TABLE GroupMembers (
    group_id INT,
    user_id INT,
    joined_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    is_admin BOOLEAN DEFAULT FALSE,
    PRIMARY KEY (group_id, user_id),
    FOREIGN KEY (group_id) REFERENCES Groups(group_id) ON DELETE CASCADE,
    FOREIGN KEY (user_id) REFERENCES Users(user_id) ON DELETE CASCADE
);
```

### 8. **Status Table** (Optional)
If you want to store users' status updates.

```sql
CREATE TABLE Status (
    status_id INT PRIMARY KEY AUTO_INCREMENT,  -- Unique identifier for each status
    user_id INT,                               -- The user who posted the status
    status_text TEXT,                          -- The content of the status
    media_url VARCHAR(255),                    -- URL to media for status (optional)
    status_type ENUM('text', 'image', 'video', 'document') NOT NULL, -- Type of status
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,  -- When the status was posted
    FOREIGN KEY (user_id) REFERENCES Users(user_id) ON DELETE CASCADE
);
```

### 9. **Notifications Table** (Optional)
This table stores notifications for users.

```sql
CREATE TABLE Notifications (
    notification_id INT PRIMARY KEY AUTO_INCREMENT,  -- Unique identifier for each notification
    user_id INT,                                     -- The user who receives the notification
    message TEXT,                                    -- Notification message
    is_read BOOLEAN DEFAULT FALSE,                   -- Whether the notification has been read
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,   -- When the notification was created
    FOREIGN KEY (user_id) REFERENCES Users(user_id) ON DELETE CASCADE
);
```

---

### Key Relationships:
1. **Users and Chats**: Users can be part of multiple chats, and chats can have multiple users (many-to-many relationship via the `UserChats` table).
2. **Chats and Messages**: Chats hold multiple messages, and each message belongs to a single chat.
3. **Messages and Media**: Messages can contain media files, which are stored separately in the `Media` table.
4. **Groups (Optional)**: Groups are essentially specialized chats with more group-specific attributes.

### Notes:
- This schema does not cover all possible features (like encryption or message statuses beyond read/unread), but it provides a foundation for a simple WhatsApp-like application.
- You may want to add indexes or constraints to improve query performance depending on the size of the dataset.
