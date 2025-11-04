# KafLite

Mini Kafka-inspired event streaming platform implemented in Go.

Stage 0: Project Scaffold.
         Set up project structure.
         Initialized Go module.

Stage 1: Basic Server & Routes.
         Integrated Gin framework.
         Created basic HTTP routes.
         Verified server response on port.
         Tested server endpoints.

Stage 2: Broker & Topics.
         Implemented Broker struct with in-memory topics (map[string]*Topic).
         Added topic creation logic using /topics endpoint (POST).
         Built publish and consume routes:
            POST /publish/:topic → Add message to topic.
            GET /consume/:topic → Fetch all messages from topic.
         Added simple unit test for topic creation.
         Verified message flow end-to-end via curl commands.