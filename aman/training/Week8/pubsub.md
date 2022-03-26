# Pub/Sub Model

Publish/subscribe messaging, or pub/sub messaging, is a form of asynchronous service-to-service communication used in serverless and microservices architectures. In a pub/sub model, any message published to a topic is immediately received by all of the subscribers to the topic.

The Pub/Sub model involves:

A publisher who sends a message
A subscriber who receives the message via a message broker

## Advantages

- Decoupled/loosely coupled components
- Greater system-wide visibility
- Real-time communication
- Ease of development
- Increased scalability & reliability
- Testability improvements

## Issues

- Unnecessary complexity in smaller systems
- Media streaming - Pub/sub is not suitable when dealing with media such as audio or video as they require smooth synchronous streaming between the host and the receiver