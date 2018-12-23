using System.Collections.Generic;

namespace MessengerServer.DbWorkers
{
    interface IDbWorker
    {
        /// <summary>
        /// Попытается открыть подключение к бд, если
        /// попытка успешна - закроет, иначе вызовет исключение.
        /// </summary>
        void CheckConnection();

        bool UserIsExist(long userId);

        void CreateUser(long userId, string login, string password);

        bool UserInDialog(long userId, long dialogId);

        void AddUserInDialog(long userId, long dialogId);

        long CreateDialog(string dialogName, ICollection<long> userIds);

        void DeleteUserFromDialog(long userId, long dialogId);

        bool GetUserId(string login, out long userId);

        ICollection<Dialog> GetAllUserDialogs(long userId);

        bool Auth(long userID, string password);

        bool CheckExist(string command);

        object ExecuteScalar(string command);

        ICollection<T> ExecuteReader<T>(string command, int columnNumb = 0);

        int ExecuteNonQuery(string command);
    }
}
