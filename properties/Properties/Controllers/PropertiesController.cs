using Microsoft.AspNetCore.Mvc;
using Properties.Data;
using Properties.Models;
using System.Collections.Generic;
using System.Linq;
using Microsoft.EntityFrameworkCore;
using System;
using System.Net.Http;
using System.Threading.Tasks;
using Microsoft.Extensions.Primitives;
using System.Text;
using System.Net;

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
        private bool Authenticate()
        {
            var request = this.Request.Headers;
            StringValues token = " ";
            if (request.TryGetValue("token", out token))
            {
                using (var client = new HttpClient())
                {
                    var req = new HttpRequestMessage(HttpMethod.Post, "http://localhost:5001/introspection/")
                    {
                        Content = new StringContent("{\"token\":\""+request["token"]+"\"}" , Encoding.UTF8, "application/json")
                     };

                    var res = client.SendAsync(req);
                    
                    var result = res.Result.StatusCode;

                    if (result.ToString() == "OK")
                    {
                        return true;
                    }
                }

            }
            return false;
        }

        [HttpGet]
        public IActionResult GetAll()
        {
            return Ok(_db.Properties.Include(p => p.Schedules).ToList());
        }

        [HttpGet("{id}")]
        public IActionResult GetById(int id)
        {
            return Ok(_db.Properties.Find(id));
        }

        [HttpPost]
        public IActionResult Create(Property property)
        {
            if (Authenticate())
            {
                _db.Properties.Add(property);
                _db.SaveChanges();
                return Ok(property);

            }
            return Unauthorized();
        }

        [HttpPost("Delete/{id}")]
        public IActionResult Delete(int id)
        {
            if (Authenticate())
            {
                Property property = _db.Properties.Find(id);
                _db.Properties.Remove(property);
                _db.SaveChanges();
                return Ok();
            }
            return Unauthorized();
        }

        [HttpPost("Update/")]
        public IActionResult Update(Property property)
        {
            if (Authenticate())
            {
                _db.Properties.Update(property);
                _db.SaveChanges();
                return Ok(property);
            }
            return Unauthorized();
        }

        [HttpGet("Schedule/{schedule_id}/{property_id}")]
        public IActionResult GetSchedule(int schedule_id, int property_id)
        {
            List<Property> properties = _db.Properties.Include(p => p.Schedules).Where(p => p.PropertyId == property_id).ToList();
            if (properties.Count() > 0)
            {
                List<Schedule> schedules = properties[0].Schedules.Where(s => s.Id == schedule_id).ToList();
                if (properties.Count() > 0) {
                    return Ok(schedules[0]);
                }          
            }
            return NotFound();
        }
    }
}
