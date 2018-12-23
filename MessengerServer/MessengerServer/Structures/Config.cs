namespace MessengerServer
{
    class Config
    {
        public string DbServer { get; set; }
        public string DbUserName { get; set; }
        public string DbUserPass { get; set; }
        public string DbName { get; set; }
        public string RMQUserName { get; set; }
        public string RMQUserPass { get; set; }
        public string VirtualHost { get; set; }
        public string HostName { get; set; }
        public QueueBind InputQueue { get; set; }
        public QueueBind OutputQueue { get; set; }
    }
}
