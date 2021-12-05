using Microsoft.AspNetCore.Mvc;
using Properties.Data;
using Properties.Models;
using System.Collections.Generic;
using System.Linq;
using Microsoft.EntityFrameworkCore;
using System;

namespace Properties.Controllers
{
    [ApiController]
    [Route("[controller]")]
    public class PropertiesController : ControllerBase
    {
        private readonly ApplicationDbContext _db;

        public PropertiesController(ApplicationDbContext db)
        {
            _db = db;
        }

        [HttpGet]
        public IEnumerable<Property> GetAll()
        {
            return _db.Properties.Include(p => p.Schedules).ToList();
        }

        [HttpGet("{id}")]
        public Property GetById(int id)
        {
            return _db.Properties.Find(id);
        }

        [HttpPost]
        public void Create(Property property)
        {
            _db.Properties.Add(property);
            _db.SaveChanges();
        }

        [HttpPost("Delete/{id}")]
        public void Delete(int id)
        {
            Property property = _db.Properties.Find(id);
            _db.Properties.Remove(property);
            _db.SaveChanges();
        }

        [HttpPost("Update/")]
        public void Update(Property property)
        {
            _db.Properties.Update(property);
            _db.SaveChanges();
        }

        [HttpGet("Schedule/{schedule_id}/{property_id}")]
        public Schedule GetSchedule(int schedule_id, int property_id)
        {
            List<Property> properties = _db.Properties.Include(p => p.Schedules).Where(p => p.PropertyId == property_id).ToList();
            if (properties.Count() > 0)
            {
                List<Schedule> schedules = properties[0].Schedules.Where(s => s.Id == schedule_id).ToList();
                if (properties.Count() > 0) {
                    return schedules[0];
                }          
            }
            throw new Exception("Schedule not found");
        }
    }
}
