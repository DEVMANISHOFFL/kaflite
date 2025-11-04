What is a Broker?

Think of the Broker as the middleman between producers and consumers.
It’s responsible for:

Storing topics

Managing message queues

Handling publish and consume operations

In Kafka, brokers are distributed.
In KafLite, we’ll start with one in-memory broker — simple but functional.





______________________________________________________________________________________________________________________________

Day 2
Create an in-memory Broker that can:
Store multiple topics
Accept messages from producers
Serve messages to consumers
______________________________________________________________________________________________________________________________

Day 3
Make KafLite’s messages persistent (survive restarts) and timestamped, plus lay groundwork for multiple consumers.

Right now pur messages live in memory.
If you stop the app — boom 💥 — everything’s gone.
Real Kafka solves this by writing every message to disk before acknowledging it.
That’s called persistence / durability — a key system-design term you’ll get asked about.
We’ll start lightweight: use simple file I/O per topic.

Plan

Add timestamps to each message
Write messages to file when published
Load existing messages when broker starts
Keep same REST APIs — they’ll just persist now

Caption:

Day 2 of building KafLite — added basic persistence 💾
Now messages survive even after server restart!
Backend devs, imagine Zomato order data surviving every crash — that’s what this step builds.
Next up: timestamps + multi-consumer support ⚡
#Golang #Kafka #SystemDesign #KafLite


✅ Day 1: Project setup + basic server
✅ Day 2: Added Broker, Topic, and 3 working endpoints
🔥 Day 3: Messages no longer vanish — you’ve added persistence