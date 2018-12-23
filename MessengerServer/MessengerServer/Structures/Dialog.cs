namespace MessengerServer
{
    class Dialog
    {
        public readonly long Id;
        public string Name { get; internal set; }

        public Dialog(long id, string name)
        {
            Id = id;
            Name = name;
        }
    }
}
