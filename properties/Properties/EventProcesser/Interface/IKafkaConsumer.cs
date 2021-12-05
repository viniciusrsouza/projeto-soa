namespace Properties.EventProcesser.Interface
{
    public interface IKafkaConsumer
    {
        string ReadMessage();
    }
}
