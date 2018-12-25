using System.Text;

namespace MessengerServer
{
    class Request
    {
        public readonly long Id;
        public string Command { get; set; }
        public string Data { get; set; }

        public Request(long id, string command, byte[] data)
        {
            Id = id;
            Command = command;
            Data = Encoding.UTF8.GetString(data);
        }
    }
}
