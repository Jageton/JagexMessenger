namespace MessengerServer
{
    class Program
    {
        static void Main(string[] args)
        {
            string file = System.IO.File.ReadAllText("Messenger.conf");
            Config conf = JsonParser.Deserialize<Config>(file);
            var mnger = new SessionManager(conf);
            mnger.Start();
        }
    }
}
