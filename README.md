# Thread-pool
This is a quick and one of the many ways to implement worker pools (aka Thread-pools).

I decided to create a thread pool using the wait group approach, an array of task and a max (thread) count for optimizing machine's resources. All this is inside of a struct type for an easy control of the pool and the task assigned to the pool.

You can run this `Thread-pool` by sending parameters to it as follows: (keep in mind that parameters must be separated by whitespaces)

    *   -parallel: Will make the calls asynchronous. If not present then it will executes calls in a synchronous way.
    *   5 (any number between 1 and 10): This number will indicate the program the max count (of threads) should use for executing the request calls.
        Note: If the number is less than 1 or greater than 10 it will automatically default to 10.
    *   From here you can set as many urls as you want to be processed.
        Note: you can set URLs in the format of 'http://google.com' or just 'google.com', in this case the system will automatically format it for you.        
    *   Program can run like this ./main -parallel 4 google.com facebook.com twitter.com github.com
 As an output the program will print out the URL of the site that is being requested and the MD5 digest (if present) on the response.  
