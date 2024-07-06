package pubsub

// Consumer is a generic consumer for different Message type
//type Consumer struct {
//	redisClient  redisDB.DB
//	subscription *redis.PubSub
//}
//
//// NewMessageConsumer creates a new instance of Consumer.
//func NewConsumer(redis redisDB.DB) *Consumer {
//	return &Consumer{
//		redisClient: redis,
//	}
//}
//
//// This Function takes queueName names in an array and uses a switch statement to perform required logic for the queues
//func (c *Consumer) Consumer(ctx context.Context, queueNames []string) {
//	for _, queueName := range queueNames {
//		switch queueName {
//		case "Test":
//			// We will handle the go routines in the custom function
//			go c.handleCustomType1Logic(ctx, queueName)
//		default:
//			log.Printf("[%s] Unsupported Message type: %+v\n", queueName, queueName)
//		}
//	}
//}
//
//func (c *Consumer) Unmarshal(message string, ptr interface{}) {
//	err := json.Unmarshal([]byte(message), ptr)
//	if err != nil {
//		log.Printf("[%s] Failed to deserialize message: %v", "User.*", err)
//	}
//}
//
//// handleCustomType1Logic initiates a goroutine to handle messages from the specified queueName.
//func (c *Consumer) handleCustomType1Logic(ctx context.Context, queueName string) {
//
//	// Create a cancellation context to gracefully stop the goroutine
//	consumerCtx, cancel := context.WithCancel(context.Background())
//	defer cancel()
//
//	log.Printf("[%s] Consumer started listening...\n", queueName)
//
//	// Subscribe to the specified Redis channel
//	c.subscription = c.redisClient.Client().Subscribe(ctx, queueName)
//	defer c.subscription.Close()
//
//	// Obtain the channel for receiving messages
//	channel := c.subscription.Channel()
//
//	for {
//		select {
//		// Check if the main context is canceled to stop the goroutine
//		case <-consumerCtx.Done():
//			log.Printf("[%s] Consumer stopped listening...\n", queueName)
//			return
//			// Listen for incoming messages on the channel
//		case msg := <-channel:
//			var messageObj interface{}
//			// Deserialize the Message payload
//			err := json.Unmarshal([]byte(msg.Payload), &messageObj)
//			if err != nil {
//				log.Printf("[%s] Failed to deserialize Message: %v", queueName, err)
//				continue
//			}
//
//			// Continue with your logic here:
//
//			fmt.Printf("[%s] Received Message: %+v\n", queueName, messageObj)
//		}
//	}
//}
