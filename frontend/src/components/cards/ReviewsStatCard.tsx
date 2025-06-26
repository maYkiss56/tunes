import { useState } from "react";
import { useProfileReviews } from "../../hooks/useProfileReviews";
import type { Review } from "../../types";
import { ReviewCard } from "./ReviewCard";
import { ReviewExpandedCard } from "./ReviewExpandCard";

export const ReviewsStatsCard = () => {
  const { reviews, loading } = useProfileReviews();
  const [expandedReview, setExpandedReview] = useState<Review | null>(null);

  const recentReviews = [...reviews].sort(
    (a, b) =>
      new Date(b.updated_at).getTime() - new Date(a.updated_at).getTime(),
  );
  return (
    <div className="w-150 h-[420px] mx-auto mt-10 mb-13 p-6 rounded-2xl bg-gradient-to-br from-gray-900 via-gray-800 to-black shadow-2xl border border-gray-700 transition-all duration-300 relative">
      <div className="flex justify-between items-center cursor-pointer">
        <div>
          <h3 className="text-xl font-bold text-white mb-1 bg-clip-text bg-gradient-to-r from-purple-400 to-pink-400">
            📝 Мои рецензии
          </h3>
          {loading ? (
            <p className="text-gray-400 text-sm">Загрузка...</p>
          ) : (
            <p className="text-white">
              <span className="text-pink-400 font-bold">{reviews.length}</span>{" "}
              {getReviewWord(reviews.length)}
            </p>
          )}
        </div>
      </div>

      <div className="mt-4 pt-4 border-t border-gray-700 max-h-[320px] overflow-y-auto pr-2 scroll scrollbar-thin scrollbar-thumb-gray-700 scrollbar-track-gray-900">
        {reviews.length === 0 ? (
          <p className="text-gray-400 text-center py-4">
            Вы еще не оставили ни одной рецензии
          </p>
        ) : (
          <>
            <div className="flex flex-col gap-y-2 ">
              {recentReviews.map((review) => (
                <ReviewCard
                  key={review.id}
                  review={review}
                  onExpand={() => setExpandedReview(review)}
                />
              ))}
            </div>

            {expandedReview && (
              <ReviewExpandedCard
                review={expandedReview}
                onClose={() => setExpandedReview(null)}
              />
            )}
          </>
        )}
      </div>
    </div>
  );
};

// Вспомогательная функция для склонения слова "рецензия"
function getReviewWord(count: number): string {
  const lastDigit = count % 10;
  const lastTwoDigits = count % 100;

  if (lastTwoDigits >= 11 && lastTwoDigits <= 19) {
    return "рецензий";
  }
  if (lastDigit === 1) {
    return "рецензия";
  }
  if (lastDigit >= 2 && lastDigit <= 4) {
    return "рецензии";
  }
  return "рецензий";
}
