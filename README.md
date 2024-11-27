### Report: Analysis of the Code Implementation and Performance

#### **Team Members**

1. Aditya Raj Kumawat (UFID: 71317106)
2. Pavan Sai Nalluri (UFID: 55592672)

---

#### **Overview**

This code simulates a platform resembling Reddit using **Proto.Actor** for actor-based concurrency and distributed systems in Go. It models core entities (e.g., users, subreddits, posts) and enables user interactions through an actor-based architecture. The system incorporates the following major actions:

1. **User Operations**: Join/leave subreddits, create posts, upvote/downvote, comment on posts or comments, send direct messages, and retrieve feeds.
2. **Subreddit Management**: Creation and user engagement simulation.
3. **Data Distribution Simulation**: Uses Zipf distribution to model subreddit popularity.

The application scales interactions and simulates activity among users to analyze the system's behavior under load.

---

#### **Key Components and Workflow**

1. **Actor-Based Concurrency**:

   - **Actors** are represented by `AccountActor` which encapsulates user-related operations.
   - Each actor processes messages asynchronously via a `Receive` method, which handles diverse user actions through message types (e.g., `JoinSubredditMsg`, `CreatePostMsg`).

2. **User and Subreddit Management**:

   - The system creates a maximum of 30,000 users (`MAX_USERS`) and manages subreddit interactions with a cap of 1,000 subreddits (`MAX_SUBREDDITS`).
   - Subreddit creation and post generation are balanced dynamically, ensuring at least 30,000 posts (`MIN_POSTS`).

3. **Message Dispatch**:

   - Random actions are selected for users, with logic to adjust priorities (e.g., forcing post creation if posts are below the threshold).

4. **Zipf Distribution**:

   - Subreddit popularity is modeled using a skewed distribution where a small number of subreddits attract the majority of members, reflecting real-world trends.

5. **Performance Simulation**:
   - Random delays (`time.Sleep`) are used to simulate staggered user interactions, mimicking real-world scenarios.

---

#### **How the Code Works**

1. **System Initialization**:

   - The actor system is created (`actor.NewActorSystem()`), and an `Engine` object maintains system state (e.g., accounts and subreddits).

2. **User Actions**:

   - Users perform random actions (e.g., joining subreddits, creating posts) chosen by a weighted random index. Actions are sent to user-specific actors via Proto.Actor's `Send` mechanism.

3. **Subreddit Interaction**:

   - Users interact with subreddits through messages like `JoinSubredditMsg`, which are processed by actors to modify the associated account's state.

4. **Performance Analysis**:
   - After the simulation, subreddit statistics are printed using the Zipf-distributed member data to evaluate subreddit popularity and engagement levels.

---

#### **Performance Analysis**

**Strengths**:

1. **Scalability**:

   - The actor model scales well, enabling concurrent operations without shared state conflicts.
   - Randomized and staggered actions simulate a high-traffic environment.

2. **Realistic Simulation**:

   - Zipf distribution provides a realistic model of subreddit popularity, highlighting the system's adaptability to real-world user behavior.

3. **Concurrency**:
   - The asynchronous messaging system minimizes blocking operations, ensuring efficient CPU utilization.

**Challenges**:

1. **High Overhead**:

   - Creating a large number of actors (one per user) and managing random actions may result in significant memory and CPU overhead.

2. **Limited Fault Handling**:

   - The code lacks mechanisms for actor failure recovery or retry logic in case of message processing errors.

3. **Simulation Bottlenecks**:

   - The use of `time.Sleep` for delays is simplistic and may not reflect complex real-world interaction patterns.

4. **Testing and Metrics**:
   - The code does not measure key performance metrics like latency, throughput, or resource utilization.

---

#### **Performance Metrics**

To fully evaluate the system's performance:

1. **Simulated Users**: Successfully handles up to 30,000 user actors and their interactions.
2. **Processing Latency**: Dependent on actor concurrency and messaging load.
3. **System Load**: Increases with the number of active users, especially during subreddit creation and post operations.

---

#### **How to run the code**

```
go run *.go
```

Output:
Logs out each and every user action/ activity

```
User User1655 got karma 0
Account created:  User1854

User User1853 joined subreddit Subreddit | 524df32a-1bd9-4b73-9077-c36524864a03
User User1853 got karma 0
User User1854 got karma 0
User User61 commented on comment in subreddit
User User61 got karma 0
User User523 joined subreddit Subreddit | 524df32a-1bd9-4b73-9077-c36524864a03
User User523 got karma 0
User User17 downvoted post in subreddit Subreddit | 524df32a-1bd9-4b73-9077-c36524864a03
User User17 got karma 15
User User896 joined subreddit Subreddit | 65af7ab7-6576-4d6b-9712-7d36e6b570ff
User User896 got karma 0
User User954 got karma 0
Account created:  User1855

User User1855 commented on comment in subreddit
User User1855 got karma 0
User User833 sent direct message to user User1738
User User833 got karma 0
User User850 is not in any subreddits
User User850 got karma 0
User User1294 created a post in subreddit Subreddit | 65af7ab7-6576-4d6b-9712-7d36e6b570ff
User User1294 got karma 0
User User842 got feed
User User842 got karma 0
Account created:  User1856

User User1788 created a post in subreddit Subreddit | 524df32a-1bd9-4b73-9077-c36524864a03
User User1788 got karma 0
User User1856 upvoted post in subreddit Subreddit | 2b053d1e-02f5-442c-9fbf-e215213b0bff
User User1856 got karma 0
Account created:  User1857

User User403 got karma 0
User User1857 commented on post in subreddit Subreddit | e35a2e07-85a8-40eb-8b15-952f5b236e33
User User1857 got karma 0
User User693 joined subreddit Subreddit | d6a6f4e2-788d-4f91-beed-0a11e98f5135
User User693 got karma 0
User User800 sent direct message to user User169
User User800 got karma -1
Account created:  User1858

User User1489 joined subreddit Subreddit | e35a2e07-85a8-40eb-8b15-952f5b236e33
User User1489 got karma 0
User User1858 upvoted post in subreddit Subreddit | 452fe59c-34c6-4fb5-a148-35ba66f83dcf
User User1858 got karma 0
Account created:  User1859

User User709 created subreddit Subreddit | 49b62059-9e0e-4668-92fa-367438d549e5
User User709 got karma 0
Account created:  User1860

User User1859 commented on comment in subreddit
User User1859 got karma 0
User User1038 is not in any subreddits
User User1038 got karma 0
```

#### **Recommendations**

1. **Optimize Actor Creation**:

   - Use a pool of reusable actors to reduce the memory overhead of creating one actor per user.

2. **Incorporate Metrics**:

   - Add logging or monitoring tools (e.g., Prometheus, Grafana) to capture performance data like message processing times and actor resource usage.

3. **Improve Fault Tolerance**:

   - Implement retry logic and error handling to manage failed message deliveries or actor crashes.

4. **Dynamic Workloads**:
   - Replace `time.Sleep` with a more realistic workload generator that dynamically adjusts action frequency based on system state.

---

#### **Conclusion**

The code provides a solid foundation for simulating a Reddit-like platform using actor-based concurrency. While it effectively demonstrates the use of Proto.Actor for managing user interactions at scale, optimization is needed for performance, fault tolerance, and metrics collection to make it suitable for production-like scenarios.
