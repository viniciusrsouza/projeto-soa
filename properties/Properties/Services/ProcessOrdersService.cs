using Microsoft.Extensions.Hosting;
using Properties.Data;
using Properties.Models;
using Properties.EventProcesser.Interface;
using System.Collections.Generic;
using System.Linq;
using System.Text.Json;
using System.Threading;
using System.Threading.Tasks;
using Microsoft.Extensions.DependencyInjection;

namespace Properties.Services
{
    public class ProcessOrdersService : BackgroundService
    {
        private readonly IKafkaConsumer _consumer;
        private readonly IServiceScopeFactory _scopeFactory;

        public ProcessOrdersService(IServiceScopeFactory scopeFactory, IKafkaConsumer consumer)
        {
            _consumer = consumer;
            _scopeFactory = scopeFactory;
        }
        protected override async Task ExecuteAsync(CancellationToken stoppingToken)
        {
            while (!stoppingToken.IsCancellationRequested)
            {
                string consumerMessage = _consumer.ReadMessage();
                if (!string.IsNullOrWhiteSpace(consumerMessage))
                {
                    OrderApprovedPayload orderApprovedPayload = JsonSerializer.Deserialize<OrderApprovedPayload>(consumerMessage);
                    using (var scope = _scopeFactory.CreateScope())
                    {
                        var db = scope.ServiceProvider.GetRequiredService<ApplicationDbContext>();
                        Property property = db.Properties.Where(p => p.OwnerId == orderApprovedPayload.Property_Owner_Id).ToList()[0];
                        List<Schedule> schedules = property.Schedules;
                        int schedule_index = property.Schedules.FindIndex(s => s.Id == orderApprovedPayload.Schedule_Id);
                        schedules[schedule_index].Tenant_Id = orderApprovedPayload.Tenant_Id;
                        schedules[schedule_index].Available = false;
                        property.Schedules = schedules;
                        db.Properties.Update(property);
                        db.SaveChanges();
                    }
                }
                else
                {
                    await Task.Delay(2000);
                }
            }
        }
    }
}
