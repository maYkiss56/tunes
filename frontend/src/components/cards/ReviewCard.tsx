import { type FC } from "react";
import type { Review } from "../../types";
import { DislikeIcon, LikeIcon } from "../ui/icons";

const BASE_URL = "http://localhost:8080/";

interface ReviewCardProps {
  review: Review;
  onExpand?: () => void;
  className?: string;
}

const ReviewCard: FC<ReviewCardProps> = ({
  review,
  onExpand,
  className = "",
}) => {
  return (
    <div
      className={`bg-gray-800 rounded-lg p-4 border border-gray-700 hover:border-purple-500 transition-colors cursor-pointer ${className}`}
      onClick={onExpand}
    >
      <div className="flex gap-4">
        <img
          src={`${BASE_URL}${review.user.avatar_url || "default-avatar.png"}`}
          alt={review.user.username}
          className="w-12 h-12 rounded-full object-cover border-2 border-purple-500 flex-shrink-0"
        />

        <div className="flex-1 min-w-0">
          <div className="flex justify-between items-center">
            <h4 className="font-semibold text-white truncate">
              {review.user.username}
            </h4>
            <span className="text-xs text-gray-400 whitespace-nowrap ml-2">
              {new Date(review.updated_at).toLocaleDateString("ru-RU", {
                day: "numeric",
                month: "long",
                year: "numeric",
              })}
            </span>
          </div>

          <div className="flex items-center gap-3 mt-3">
            <img
              src={`${BASE_URL}${review.song.image_url || "default-song.png"}`}
              alt={review.song.title}
              className="w-10 h-10 rounded object-cover"
            />

            <div className="min-w-0">
              <p className="text-sm font-medium text-white truncate">
                {review.song.title}
              </p>
              <div className="flex items-center gap-1 mt-1">
                {review.is_like ? (
                  <>
                    <LikeIcon className="text-green-500 w-4 h-4" />
                    <span className="text-xs text-green-500">Лайк</span>
                  </>
                ) : (
                  <>
                    <DislikeIcon className="text-red-500 w-4 h-4" />
                    <span className="text-xs text-red-500">Дизлайк</span>
                  </>
                )}
              </div>
            </div>
          </div>

          {review.body && (
            <p className="mt-3 text-gray-300 text-sm line-clamp-2">
              {review.body}
            </p>
          )}
        </div>
      </div>
    </div>
  );
};

export { ReviewCard };
