import { useEffect, useRef, useState, type ChangeEvent, type FC } from "react";
import { Modal } from "../ui/Modal";

const ProfileCard: FC = () => {
  const [email, setEmail] = useState("");
  const [username, setUsername] = useState("");
  const [imagePreview, setImagePreview] = useState("");
  const [avatar, setAvatar] = useState<File | null>(null);
  const [showPasswordModal, setShowPasswordModal] = useState(false);
  const [currentPassword, setCurrentPassword] = useState("");
  const [newPassword, setNewPassword] = useState("");
  const [confirmPassword, setConfirmPassword] = useState("");
  const [passwordError, setPasswordError] = useState("");

  const fileInputRef = useRef<HTMLInputElement>(null);

  useEffect(() => {
    const fetchProfile = async () => {
      const res = await fetch("http://localhost:8080/api/profile", {
        credentials: "include",
      });

      if (res.ok) {
        const data = await res.json();
        setEmail(data.email);
        setUsername(data.username);

        if (data.avatar_url) {
          const fullUrl = data.avatar_url.startsWith("http")
            ? data.avatar_url
            : `http://localhost:8080/${data.avatar_url}`;
          setImagePreview(fullUrl);
          setAvatar(null);
        }
      }
    };

    fetchProfile();
  }, []);

  const handleFileChange = (e: ChangeEvent<HTMLInputElement>) => {
    const file = e.target.files?.[0];
    if (file) {
      setAvatar(file);
      const reader = new FileReader();
      reader.onload = () => {
        setImagePreview(reader.result as string);
      };
      reader.readAsDataURL(file);
    }
  };

  const handleAvatarSubmit = async () => {
    if (!avatar) return;

    const formData = new FormData();
    formData.append("avatar", avatar);

    const res = await fetch("http://localhost:8080/api/profile/avatar", {
      method: "PUT",
      body: formData,
      credentials: "include",
    });

    if (res.ok) {
      const data = await res.json();
      const fullUrl = data.avatar_url.startsWith("http")
        ? data.avatar_url
        : `http://localhost:8080/${data.avatar_url}`;
      setImagePreview(fullUrl);
      setAvatar(null);
    }
  };

  const handlePasswordSubmit = async () => {
    if (newPassword !== confirmPassword) {
      setPasswordError("Пароли не совпадают");
      return;
    }

    if (newPassword.length < 8) {
      setPasswordError("Пароль должен содержать минимум 8 символов");
      return;
    }

    try {
      const res = await fetch("http://localhost:8080/api/profile/password", {
        method: "PUT",
        headers: {
          "Content-Type": "application/json",
        },
        credentials: "include",
        body: JSON.stringify({
          old_password: currentPassword,
          new_password: newPassword,
        }),
      });

      if (res.ok) {
        setShowPasswordModal(false);
        setCurrentPassword("");
        setNewPassword("");
        setConfirmPassword("");
        setPasswordError("");
        // Можно добавить уведомление об успешном изменении
      } else {
        const error = await res.json();
        setPasswordError(error.message || "Не удалось изменить пароль");
      }
    } catch (err) {
      setPasswordError("Произошла ошибка при изменении пароля");
    }
  };

  return (
    <>
      <div className="max-w-md mx-auto mt-10 p-6 rounded-2xl bg-gradient-to-br from-gray-900 via-gray-800 to-black shadow-2xl border border-gray-700 mb-13">
        <h2 className="text-2xl font-bold text-white mb-4 bg-clip-text bg-gradient-to-r from-purple-400 to-pink-400">
          👤 Профиль пользователя
        </h2>

        {imagePreview && (
          <div className="mb-4 flex justify-center">
            <img
              src={imagePreview}
              alt="Аватар пользователя"
              className="w-24 h-24 rounded-full border-2 border-pink-500 object-cover transition-opacity duration-300"
            />
          </div>
        )}

        <div className="space-y-2 text-gray-300 text-sm sm:text-base mb-6">
          <div className="flex flex-col">
            <span className="text-gray-400 font-medium">Имя пользователя:</span>
            <span className="text-white">{username}</span>
          </div>
          <div className="flex flex-col">
            <span className="text-gray-400 font-medium">Email:</span>
            <span className="text-white">{email}</span>
          </div>
        </div>

        <div className="space-y-3">
          <input
            ref={fileInputRef}
            type="file"
            accept="image/*"
            onChange={handleFileChange}
            className="hidden"
          />

          <button
            onClick={() => fileInputRef.current?.click()}
            className="w-full px-4 py-2 bg-pink-600 hover:bg-pink-700 text-white rounded-xl transition duration-200"
            type="button"
          >
            {avatar ? "Изменить изображение" : "Загрузить аватар"}
          </button>

          {avatar && (
            <button
              onClick={handleAvatarSubmit}
              className="w-full px-4 py-2 bg-purple-600 hover:bg-purple-700 text-white rounded-xl transition duration-200"
              type="button"
            >
              Сохранить аватар
            </button>
          )}

          <button
            onClick={() => setShowPasswordModal(true)}
            className="w-full px-4 py-2 bg-blue-600 hover:bg-blue-700 text-white rounded-xl transition duration-200"
            type="button"
          >
            Изменить пароль
          </button>
        </div>
      </div>

      {/* Модальное окно для смены пароля */}
      <Modal
        isOpen={showPasswordModal}
        onClose={() => setShowPasswordModal(false)}
      >
        <div className="p-6 bg-gray-800 rounded-lg">
          <h3 className="text-xl font-bold text-white mb-4">
            Изменение пароля
          </h3>

          <div className="space-y-4">
            <div>
              <label className="block text-gray-300 mb-1">Текущий пароль</label>
              <input
                type="password"
                value={currentPassword}
                onChange={(e) => setCurrentPassword(e.target.value)}
                className="w-full px-4 py-2 bg-gray-700 border border-gray-600 rounded-md text-white"
              />
            </div>

            <div>
              <label className="block text-gray-300 mb-1">Новый пароль</label>
              <input
                type="password"
                value={newPassword}
                onChange={(e) => setNewPassword(e.target.value)}
                className="w-full px-4 py-2 bg-gray-700 border border-gray-600 rounded-md text-white"
              />
            </div>

            <div>
              <label className="block text-gray-300 mb-1">
                Подтвердите пароль
              </label>
              <input
                type="password"
                value={confirmPassword}
                onChange={(e) => setConfirmPassword(e.target.value)}
                className="w-full px-4 py-2 bg-gray-700 border border-gray-600 rounded-md text-white"
              />
            </div>

            {passwordError && (
              <div className="text-red-400 text-sm">{passwordError}</div>
            )}

            <div className="flex justify-end space-x-3 pt-2">
              <button
                onClick={() => setShowPasswordModal(false)}
                className="px-4 py-2 text-gray-300 hover:text-white"
              >
                Отмена
              </button>
              <button
                onClick={handlePasswordSubmit}
                className="px-4 py-2 bg-blue-600 hover:bg-blue-700 text-white rounded-md"
              >
                Сохранить
              </button>
            </div>
          </div>
        </div>
      </Modal>
    </>
  );
};

export { ProfileCard };
