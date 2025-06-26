import { type FC, useState, useEffect } from "react";
import { Button } from "../ui/Button";
import { useAuth } from "../../context/AuthContext";
import type { Review, Track } from "../../types";
import { CloseIcon, DislikeIcon, LikeIcon } from "../ui/icons";

interface ModelWindowProps {
  track: Track;
  onClose: () => void;
}

const ModelWindow: FC<ModelWindowProps> = ({ track, onClose }) => {
  const { user } = useAuth();
  const [authWarning, setAuthWarning] = useState(false);
  const [artistName, setArtistName] = useState<string>("");
  const [albumTitle, setAlbumTitle] = useState<string>("");
  const [reviews, setReviews] = useState<Review[]>([]);
  const [reviewText, setReviewText] = useState("");
  const [isLike, setIsLike] = useState<boolean | null>(null);
  const [editingId, setEditingId] = useState<number | null>(null);
  const [showForm, setShowForm] = useState(false);
  const [isNewReview, setIsNewReview] = useState(false);

  useEffect(() => {
    fetchMeta();
    fetchReviews();
  }, [track]);

  const fetchMeta = async () => {
    try {
      const [artistRes, albumRes] = await Promise.all([
        fetch(`http://localhost:8080/api/artists/${track.artist.id}`),
        fetch(`http://localhost:8080/api/albums/${track.album?.id}`),
      ]);
      const artist = await artistRes.json();
      const album = await albumRes.json();
      setArtistName(artist.nickname || "");
      setAlbumTitle(album.title || "");
    } catch (e) {
      console.error("Ошибка загрузки метаданных:", e);
    }
  };

  const fetchReviews = async () => {
    try {
      const res = await fetch("http://localhost:8080/api/reviews");
      const data = await res.json();
      setReviews(data);
    } catch (e) {
      console.error("Ошибка загрузки рецензий:", e);
    }
  };

  const submitReview = async () => {
    if (!reviewText.trim() || isLike === null) return;

    const payload = {
      body: reviewText,
      is_like: isLike,
      song_id: track.id,
    };

    const method = editingId ? "PATCH" : "POST";
    const url = editingId
      ? `http://localhost:8080/api/reviews/${editingId}`
      : "http://localhost:8080/api/reviews";

    try {
      await fetch(url, {
        method,
        headers: { "Content-Type": "application/json" },
        credentials: "include",
        body: JSON.stringify(payload),
      });

      fetchReviews();
      cancelReview();
    } catch (e) {
      console.error("Ошибка при отправке:", e);
    }
  };

  const cancelReview = () => {
    setShowForm(false);
    setReviewText("");
    setIsLike(null);
    setEditingId(null);
    setIsNewReview(false);
  };

  const editReview = (r: Review) => {
    setShowForm(true);
    setReviewText(r.body);
    setIsLike(r.is_like);
    setEditingId(r.id);
    setIsNewReview(false);
  };

  const startNewReview = () => {
    if (!user) {
      setAuthWarning(true);
      return;
    }
    setAuthWarning(false);
    setShowForm(true);
    setIsNewReview(true);
    setEditingId(null);
    setReviewText("");
    setIsLike(null);
  };

  const deleteReview = async (id: number) => {
    try {
      await fetch(`http://localhost:8080/api/reviews/${id}`, {
        method: "DELETE",
        credentials: "include",
      });
      fetchReviews();
    } catch (e) {
      console.error("Ошибка удаления:", e);
    }
  };

  const formatDate = (date: string) =>
    new Date(date).toLocaleDateString("ru-RU", {
      day: "numeric",
      month: "long",
      year: "numeric",
    });

  return (
    <div className="fixed inset-0 bg-black/70 z-50 flex items-start justify-center overflow-y-auto p-4 sm:p-10">
      <div className="bg-gray-900 rounded-2xl w-full max-w-3xl p-6 relative text-white shadow-lg border border-gray-700">
        <button
          onClick={onClose}
          className="absolute top-4 right-4 text-gray-400 hover:text-white text-2xl font-bold cursor-pointer transition"
        >
          <CloseIcon />
        </button>

        <div className="flex flex-col md:flex-row gap-6">
          <div className="flex-shrink-0">
            <img
              src={`http://localhost:8080/${track.image_url}`}
              alt={track.title}
              className="w-full h-64 object-cover rounded-xl"
            />
          </div>

          <div>
            <h2 className="text-2xl font-bold mb-2">{track.title}</h2>
            <div className="mb-4 text-sm space-y-1 text-gray-300">
              <p>
                <strong className="text-gray-400">Дата релиза:</strong>{" "}
                {formatDate(track.release_date)}
              </p>
              <p>
                <strong className="text-gray-400">Альбом:</strong> {albumTitle}
              </p>
              <p>
                <strong className="text-gray-400">Артист:</strong> {artistName}
              </p>
            </div>
          </div>
        </div>

        <div className="mt-8">
          <div className="flex justify-between items-center mb-4">
            <h3 className="text-xl font-semibold">Рецензии</h3>
            {!showForm && (
              <Button
                className="bg-purple-600 hover:bg-purple-700 text-white px-4 py-2 transition"
                onClick={startNewReview}
              >
                Новая рецензия
              </Button>
            )}
          </div>

          {authWarning && (
            <div className="text-red-400 text-sm mb-4">
              Пожалуйста, авторизуйтесь, чтобы оставить рецензию.
            </div>
          )}

          {showForm ? (
            <div className="space-y-4 bg-gray-800 p-4 rounded-lg border border-gray-700 mb-6">
              <textarea
                value={reviewText}
                onChange={(e) => setReviewText(e.target.value)}
                className="w-full bg-gray-900 border border-gray-700 rounded-lg p-3 resize-none text-sm text-white focus:outline-none focus:ring-2 focus:ring-purple-500 focus:border-transparent"
                rows={4}
                placeholder="Напишите вашу рецензию..."
              />

              <div>
                <label className="block text-sm font-medium mb-2 text-gray-300">
                  Ваша оценка
                </label>
                <div className="flex gap-4">
                  <Button
                    className={`flex-1 flex items-center justify-center gap-2 py-2 ${
                      isLike === true
                        ? "bg-green-600 hover:bg-green-700 text-white"
                        : "bg-gray-700 hover:bg-gray-600 text-gray-300"
                    }`}
                    onClick={() => setIsLike(true)}
                  >
                    <LikeIcon className="w-5 h-5" />
                    Лайк
                  </Button>
                  <Button
                    className={`flex-1 flex items-center justify-center gap-2 py-2 ${
                      isLike === false
                        ? "bg-red-600 hover:bg-red-700 text-white"
                        : "bg-gray-700 hover:bg-gray-600 text-gray-300"
                    }`}
                    onClick={() => setIsLike(false)}
                  >
                    <DislikeIcon className="w-5 h-5" />
                    Дизлайк
                  </Button>
                </div>
              </div>

              <div className="flex gap-3 pt-2">
                <Button
                  className="flex-1 bg-green-600 hover:bg-green-700 text-white"
                  onClick={submitReview}
                  disabled={isLike === null}
                >
                  {editingId ? "Сохранить" : "Опубликовать"}
                </Button>
                <Button
                  className="flex-1 bg-gray-700 hover:bg-gray-600 text-gray-300"
                  onClick={cancelReview}
                >
                  Отмена
                </Button>
              </div>
            </div>
          ) : (
            <div className="space-y-4 mb-6 max-h-96 overflow-y-auto pr-2 scrollbar-thin scrollbar-thumb-gray-700 scrollbar-track-gray-900">
              {reviews
                .filter((r) => r.song.id === track.id)
                .map((r) => (
                  <div
                    key={r.id}
                    className="bg-gray-800 p-4 rounded-lg border border-gray-700"
                  >
                    <div className="flex gap-3">
                      <div className="flex-shrink-0">
                        <img
                          src={`http://localhost:8080/${r.user.avatar_url || "default-avatar.png"}`}
                          alt={r.user.username}
                          className="w-10 h-10 rounded-full object-cover border-2 border-purple-600"
                        />
                      </div>

                      <div className="flex-1">
                        <div className="flex items-center justify-between">
                          <div className="flex items-center gap-2">
                            <span className="font-medium text-white">
                              {r.user.username}
                            </span>
                            <span
                              className={`inline-flex items-center px-2 py-0.5 rounded-full text-xs font-medium ${
                                r.is_like
                                  ? "bg-green-900/50 text-green-400"
                                  : "bg-red-900/50 text-red-400"
                              }`}
                            >
                              {r.is_like ? (
                                <>
                                  <LikeIcon className="w-3 h-3 mr-1" />
                                  Лайк
                                </>
                              ) : (
                                <>
                                  <DislikeIcon className="w-3 h-3 mr-1" />
                                  Дизлайк
                                </>
                              )}
                            </span>
                          </div>

                          {r.user.id === user?.id && (
                            <div className="flex gap-2">
                              <button
                                onClick={() => editReview(r)}
                                className="text-gray-400 hover:text-blue-400 text-sm transition cursor-pointer"
                              >
                                Редактировать
                              </button>
                              <button
                                onClick={() => deleteReview(r.id)}
                                className="text-gray-400 hover:text-red-400 text-sm transition cursor-pointer"
                              >
                                Удалить
                              </button>
                            </div>
                          )}
                        </div>

                        <p className="mt-2 text-gray-300 text-sm">{r.body}</p>
                        <p className="mt-2 text-xs text-gray-500">
                          {formatDate(r.updated_at)}
                        </p>
                      </div>
                    </div>
                  </div>
                ))}
            </div>
          )}
        </div>
      </div>
    </div>
  );
};

export default ModelWindow;
