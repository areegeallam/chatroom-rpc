# RPC Chatroom — Assignment 04

**Student:** Areege Allam  
**Course:** Distributed Systems — Level 4 (Fall 2025)

---

## 🧠 Project Description
This project implements a simple **chatroom using Go RPC (Remote Procedure Calls)**.  
The server stores all messages in memory and returns the chat history whenever a new message is sent.

---

## ⚙️ How It Works

### 🖥️ Server
- Listens on port `1234`
- Handles two RPC methods:
  - `SendMessage(message string, reply *[]string)`
  - `GetHistory(dummy string, reply *[]string)`
- Stores all chat messages in a list

### 💻 Client
- Connects to the server using `net/rpc.Dial("tcp", "localhost:1234")`
- Reads full-line input using `bufio.NewReader`
- Sends messages and receives chat history
- Runs continuously until the user types `"exit"`

---

## 🚀 How to Run

```bash
# Run the server
go run server.go

# In another terminal
go run client.go
```

You can open multiple clients for group chatting!

---

## 💬 Example Output

### 🧩 Client 1:
```
Enter message: Hello everyone!
📜 Chat history:
• Hello everyone!
```

### 🧩 Client 2:
```
Enter message: Hi there!
📜 Chat history:
• Hello everyone!
• Hi there!
```

When typing:
```
exit
```
the client disconnects safely.

---

## 🚨 Error Handling
If the server goes down while a client is connected, the client prints:
```
Connection error: dialing failed
```
and exits gracefully.

---

## 🎥 Demo Video

📹 [Click here to watch the running project](https://drive.google.com/file/d/1rhfL70BO6taj5nBGant9bVSYekAUktCB/view?usp=drive_link)

(Example:  
`https://drive.google.com/file/d/1abcXYZ123/view?usp=sharing`  
)

---

## 📄 Notes
- The server stores messages **in memory only** (not in a file or database).  
- When the server restarts, the chat history is cleared.  
- The client reads **full-line input** using `bufio.NewReader` so spaces in messages are allowed.  

---

## 🧱 Technologies Used
- Go 1.25.3  
- net/rpc package  
- Windows Terminal / VS Code  

---

## ✍️ Author
Developed by **Areege Allam**  
Tanta University — Faculty of Engineering  
Distributed Systems, Fall 2025

---

## 🧾 Submission
- GitHub Repo: [https://github.com/areegeallam/chatroom-rpc](https://github.com/areegeallam/chatroom-rpc)  
- Demo Video: (add your Google Drive link above)
