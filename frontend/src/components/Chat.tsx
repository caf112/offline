import { useState } from "react";
import { useWebSocket } from "../hooks/useWebSocket";

export default function Chat() {
  const { messages, sendMessage } = useWebSocket("ws://localhost:8080/ws");
  const [input, setInput] = useState("");

  console.log(messages)

  return (
    <div className="p-6 max-w-md mx-auto">
      <h2 className="text-2xl font-semibold mb-4">offline. chat</h2>

      <div className="border rounded p-3 h-64 overflow-y-auto bg-gray-50 mb-3">
        {messages.map((msg, i) => (
          <div key={i} className="text-gray-800 mb-1">
            {msg}
          </div>
        ))}
      </div>

      <div className="flex gap-2">
        <input
          className="flex-1 border px-2 py-1 rounded"
          value={input}
          placeholder="メッセージを入力..."
          onChange={(e) => setInput(e.target.value)}
        />
        <button
          className="bg-blue-500 text-white px-4 py-1 rounded"
          onClick={() => {
            sendMessage(input);
            setInput("");
          }}
        >
          送信
        </button>
      </div>
    </div>
  );
}
