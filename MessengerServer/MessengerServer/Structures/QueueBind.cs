namespace MessengerServer
{
    struct QueueBind
    {
        public string Name { get; set; }
        public string Exchange { get; set; }
        public string RoatKey { get; set; }

        public QueueBind(string name, string exchange, string roatKey)
        {
            Name = name;
            Exchange = exchange;
            RoatKey = roatKey;
        }
    }
}
