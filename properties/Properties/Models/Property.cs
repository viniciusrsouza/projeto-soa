using System.Collections.Generic;
using System.ComponentModel.DataAnnotations;

namespace Properties.Models
{
    public class Property
    {
        [Key]
        public int PropertyId { get; set; }

        [Required]
        public int OwnerId { get; set; }

        public List<Schedule> Schedules { get; set; }

        [Required]
        public string Address { get; set; }

        public string Description { get; set; }
    }
}
