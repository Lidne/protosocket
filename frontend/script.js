// Load the protobuf definition first
protobuf.load("../proto/mymessage.proto", function (err, root) {
  if (err) {
    console.error("Failed to load proto:", err);
    return;
  }

  // Retrieve the type of message I want to decode
  const MyMessage = root.lookupType("MyMessage");

  const socket = new WebSocket("ws://localhost:8000/ws");
  socket.binaryType = "arraybuffer";

  // Connection opened
  socket.addEventListener("open", (event) => {
    console.log("Connected to WebSocket");
  });

  // Listen for messages
  socket.addEventListener("message", function (event) {
    try {
      // event.data is an ArrayBuffer because binaryType is "arraybuffer"
      const buffer = new Uint8Array(event.data);

      // Decode the message
      const message = MyMessage.decode(buffer);

      // message now contains an object with the properties specified in the .proto file
      console.log(
        `Received: ID=${message.id}, Title="${message.title}", Text="${message.text}"`
      );
      console.log(message);
    } catch (e) {
      console.error("Failed to decode message:", e);
    }
  });

  // Listen for Connection closure
  socket.addEventListener("close", function () {
    console.log("Connection closed");
  });

  socket.addEventListener("error", function (err) {
    console.error("WebSocket error:", err);
  });
});
