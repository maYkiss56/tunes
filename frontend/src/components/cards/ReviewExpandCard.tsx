import { type FC } from "react";
import type { Review } from "../../types";
import { CloseIcon, DislikeIcon, LikeIcon } from "../ui/icons";

const BASE_URL = "http://localhost:8080/";

interface ReviewExpandedCardProps {
  review: Review;
  onClose: () => void;
}

const ReviewExpandedCard: FC<ReviewExpandedCardProps> = ({
  review,
  onClose,
}) => {
  return (
    <div className="fixed inset-0 bg-black/70 z-50 flex items-center justify-center p-4">
      <div className="bg-gray-800 rounded-lg max-w-2xl w-full relative border border-gray-700 overflow-hidden">
        <div className="flex justify-between items-center p-4 border-b border-gray-700 bg-gray-900/50">
          <h3 className="text-xl font-bold text-white">Рецензия</h3>
          <button
            onClick={onClose}
            className="text-gray-400 hover:text-white transition-colors cursor-pointer"
            aria-label="Закрыть"
          >
            <CloseIcon />
          </button>
        </div>

        <div className="p-6 overflow-y-auto max-h-[80vh]">
          <div className="flex items-start gap-4 mb-6">
            <img
              src={`${BASE_URL}${review.user.avatar_url || "default-avatar.png"}`}
              alt={review.user.username}
              className="w-14 h-14 rounded-full object-cover border-2 border-purple-500 flex-shrink-0"
            />
            <div className="flex-1 min-w-0">
              <div className="flex flex-col sm:flex-row sm:justify-between sm:items-start gap-2">
                <div>
                  <h3 className="font-bold text-white truncate">
                    {review.user.username}
                  </h3>
                  <div className="flex items-center gap-2 mt-1">
                    {review.is_like ? (
                      <LikeIcon className="text-green-500 w-4 h-4 flex-shrink-0" />
                    ) : (
                      <DislikeIcon className="text-red-500 w-4 h-4 flex-shrink-0" />
                    )}
                    <span className="text-sm text-gray-400">
                      {review.is_like ? "Лайк" : "Дизлайк"}
                    </span>
                  </div>
                </div>
                <span className="text-sm text-gray-500 whitespace-nowrap">
                  {new Date(review.updated_at).toLocaleDateString("ru-RU", {
                    day: "numeric",
                    month: "long",
                    year: "numeric",
                  })}
                </span>
              </div>
            </div>
          </div>

          <div className="mb-6 bg-gray-900/50 p-4 rounded-lg">
            <p className="text-gray-300 whitespace-pre-line break-words">
              {review.body}
            </p>
          </div>

          <div className="flex items-center gap-3 bg-gray-700/50 p-3 rounded-lg mb-6 hover:bg-gray-700/70 transition-colors">
            <img
              src={`${BASE_URL}${review.song.image_url}`}
              alt={review.song.title}
              className="w-16 h-16 rounded-md object-cover flex-shrink-0"
            />
            <div className="min-w-0">
              <h4 className="font-semibold text-white truncate">
                {review.song.title}
              </h4>
              <p className="text-sm text-gray-400 truncate">
                {review.song.full_title}
              </p>
            </div>
          </div>
        </div>
      </div>
    </div>
  );
};

export { ReviewExpandedCard };
