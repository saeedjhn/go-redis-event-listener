Redis Pub-Sub, short for Publish-Subscribe, is a messaging pattern where message senders (publishers) distribute messages to multiple
recipients (subscribers) without the need for direct communication between them using different channels that will listen to events.

### Where usage message broker
Message brokers are widely used in distributed systems and microservices architectures to enable asynchronous communication between different services or components. They decouple the sender and receiver, allowing them to communicate without being directly connected. This helps improve scalability, reliability, and flexibility of applications. Here are some common use cases for message brokers:

1. Microservices Communication
In a microservices architecture, different services often need to communicate with each other. A message broker can facilitate this communication by allowing services to publish messages to topics or channels that other services can subscribe to. This decouples services, allowing them to scale independently and reducing direct dependencies.

2. Event-Driven Architecture
Message brokers are essential in event-driven architectures where components respond to events rather than direct calls. For example, an e-commerce application might publish events such as "order placed", "payment received", or "item shipped". Different services can subscribe to these events and perform actions like updating inventory, notifying the customer, or triggering shipping processes.

3. Task Queues
Message brokers can be used to implement task queues. For instance, a web application might need to perform time-consuming tasks like image processing or sending emails. These tasks can be added to a queue, and worker services can consume and process them asynchronously, improving the application's responsiveness.

4. Load Balancing and Scalability
Message brokers can distribute messages across multiple consumers, helping to balance the load and improve the scalability of applications. For example, a logging system can publish log messages to a broker, and multiple log processing services can consume and process these messages concurrently.

5. Data Streaming
Message brokers can handle real-time data streams. For instance, in IoT applications, sensors might send a continuous stream of data to a broker, which can then be processed by different services for monitoring, alerting, and analysis.

6. Decoupling Components
Message brokers help decouple components, making it easier to manage dependencies and evolve the system. For instance, a payment service might publish a "payment completed" event, which different services (like order processing, inventory management, and notifications) can handle independently.

7. Reliable Message Delivery
Message brokers can ensure reliable delivery of messages, even in the case of network failures or service crashes. They often provide mechanisms for message persistence, acknowledgments, and retries, ensuring that messages are not lost.

more:

Understanding MQTT Quality of Service or also known as MQTT QoS
https://cedalo.com/blog/understanding-mqtt-qos/

Message Brokers: Key Models, Use Cases & Tools Simplified 101
https://hevodata.com/learn/message-brokers/

Dissecting Message Queues
https://bravenewgeek.com/dissecting-message-queues/

https://developer.mozilla.org/en-US/docs/Glossary/Idempotent
idempotent

Choosing The Right Message Queue Technology For Your Notification System
https://dev.to/nikl/choosing-the-right-message-queue-technology-for-your-notification-system-2mji

How choose message queue technology selection
https://blog.iron.io/how-to-choose-message-queue-technology-selection/