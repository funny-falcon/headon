using System;
using System.Diagnostics;
using System.Threading.Tasks;

namespace ConsoleApplication
{
    public class Program
    {
        public static void Main(string[] args)
        {
            if (args == null || args.Length != 1)
            {
                Console.WriteLine("invalid argument");
                return;
            }

            var taskCount = Convert.ToInt64(args[0]);

            Console.WriteLine("Task to execute: {0}", taskCount);

            var sw = Stopwatch.StartNew();

            var tasks = new Task<uint>[taskCount];

            for (int i = 0; i < tasks.Length; i++)
            {
                 int j = i;
                 tasks[i] = Task.Run(()=>Work(j));
            }

            Task.WaitAll(tasks);

	    uint h = 0;
	    for (int i = 0; i < tasks.Length; i++) {
		 h ^= tasks[i].Result;
	    }
            // Parallel.For(0, taskCount, (i)=> {
            //     var t = string.Format("Task {0} done!", i);
            // });            

            sw.Stop();
            Console.WriteLine("{0} in {1}, hash = 0x{2:x}", taskCount, sw.Elapsed, h);
        }

        private static async Task<uint> Work(int i){
            var t = string.Format("Task {0} done!", i);
	    await Task.Delay(TimeSpan.FromMilliseconds(0.1));
	    return hash(t);
        }

	private static uint hash(string s) {
	    uint h = 0xcafe;
	    int l = s.Length;
	    for (int i = 0; i < l; i++) {
		    h ^= s[i];
		    h *= 0x11f;
	    }
	    return h;
	}
    }
}
