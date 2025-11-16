import { useEffect, useState } from "react";

export const useWebSocket = (url: string) => {
  const [messages, setMessages] = useState<string[]>([]);
  const [socket, setSocket] = useState<WebSocket | null>(null);

  useEffect(() => {
    const ws = new WebSocket(url);
    setSocket(ws);

    ws.onmessage = (event) => {
      setMessages((prev) => [...prev, event.data]);
    };

    ws.onclose = () => {
      console.log("❌ WebSocket closed");
    };

    return () => ws.close();
  }, [url]);

  const sendMessage = (msg: string) => {
    if (socket?.readyState === WebSocket.OPEN) {
      socket.send(msg);
    }
  };

  return { messages, sendMessage };
};
