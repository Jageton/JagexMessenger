namespace MessengerServer
{
    class Request
    {
        public readonly long Id;
        public string Command { get; internal set; }
        public string Data { get; internal set; }

        public Request(long id, string cmd, string data)
        {
            Id = id;
            Command = cmd;
            Data = data;
        }
    }
}
