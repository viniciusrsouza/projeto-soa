using Properties.EventProcesser.Interface;
using System;
using System.Collections.Generic;
using Confluent.Kafka;

namespace Properties.EventProcesser.Kafka
{
    public class KafkaConsumer : IKafkaConsumer
    {
        private readonly IConsumer<string, string> _consumer;
        public KafkaConsumer()
        {
            _consumer = KafkaUtils.CreateConsumer("localhost:9092", new List<string> { "domain_orderservice_schedule_approved_0:2:1" });
        }

        public string ReadMessage()
        {
            var consumeResult = _consumer.Consume(TimeSpan.FromSeconds(0.5));
            if (consumeResult == null || string.IsNullOrWhiteSpace(consumeResult.Message.Value))
            {
                return string.Empty;
            }
            else
            {
                return consumeResult.Message.Value;
            }
        }
    }
}
