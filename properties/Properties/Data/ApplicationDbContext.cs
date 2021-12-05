using Properties.Models;
using Microsoft.EntityFrameworkCore;

namespace Properties.Data
{
    public class ApplicationDbContext : DbContext
    {
        public ApplicationDbContext(DbContextOptions<ApplicationDbContext> options) : base(options)
        {

            this.ChangeTracker.LazyLoadingEnabled = false;
        }
        public DbSet<Property> Properties { get; set; }
    }
}
