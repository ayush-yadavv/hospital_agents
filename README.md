### Prerequisites

- Go 1.23.4 or higher
- Gemini API key from https://aistudio.google.com
- `.env` file with GOOGLE_GENAI_API_KEY

### Installation

1. Clone the repository:

   ```bash
   git clone https://github.com/xprilion/go-tiny-agents.git
   cd go-tiny-agents
   ```

2. Create a `.env` file:

   ```bash
   GOOGLE_GENAI_API_KEY=your_api_key_here
   ```

3. Install dependencies:

   ```bash
   go mod download
   ```

4. Run the application:
   ```bash
   go run .
   ```

## 📝 Usage

### API Endpoints

#### POST /message

Send a message to the agents.

**Request Body:**

```json
{
  "message": "Hello, how are you?"
}
```

**Response:**

```json
[
  {
    "name": "Michael Scott",
    "message": "Hello, how are you?"
  }
]
```
