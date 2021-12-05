using Confluent.Kafka;
using System.Collections.Generic;

namespace Properties.EventProcesser.Kafka
{
    public static class KafkaUtils
    {
        public static IConsumer<string, string> CreateConsumer(string brokerList, List<string> topics)
        {
            var config = new ConsumerConfig
            {
                BootstrapServers = brokerList,
                GroupId = "appOrder",
                // Note: The AutoOffsetReset property determines the start offset in the event
                // there are not yet any committed offsets for the consumer group for the
                // topic/partitions of interest. By default, offsets are committed
                // automatically, so in this example, consumption will only start from the
                // earliest message in the topic 'my-topic' the first time you run the program.
                AutoOffsetReset = AutoOffsetReset.Earliest
            };
            var consumer = new ConsumerBuilder<string, string>(config).Build();
            consumer.Subscribe(topics);
            return consumer;
        }
    
    }
}
