using System;
using System.ComponentModel.DataAnnotations;

namespace Properties.Models
{
    public class Schedule
    {
        [Key]
        public int Id { get; set; }

        [Required]
        public DateTime Start { get; set; }

        [Required]
        public DateTime End { get; set; }

        [Required]
        public bool Available { get; set; }

        [Required]
        public int Owner_Id { get; set; }

        public int Tenant_Id { get; set; }

    }
}
