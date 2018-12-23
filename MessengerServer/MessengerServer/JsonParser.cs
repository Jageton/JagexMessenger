using System.Collections.Generic;
using Newtonsoft.Json;

namespace MessengerServer
{
    static class JsonParser
    {
        public static KeyValuePair<K, V> JsonToKeyValuePair<K, V>(string str)
        {
            return JsonConvert.DeserializeObject
                <KeyValuePair<K, V>>(str);
        }

        public static string KeyValuePairToJson<K, V>(KeyValuePair<K, V> pair)
        {
            return JsonConvert.SerializeObject(pair);
        }

        public static string Serialize(object type)
        {
            return JsonConvert.SerializeObject(type);
        }

        public static string SerializeException(string exception)
        {
            return Serialize(new { Exception = exception });
        }

        public static T Deserialize<T>(string obj)
        {
            return JsonConvert.DeserializeObject<T>(obj);
        }

        public static T DeserializeAnonymousType<T>(string obj, T type)
        {
            return JsonConvert.DeserializeAnonymousType(obj, type);
        }
    }
}
