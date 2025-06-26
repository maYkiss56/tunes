import { useState } from "react";
import { useRecentReviews } from "../../hooks/useRecentReviews";
import type { Review } from "../../types";
import { ReviewCard } from "../cards/ReviewCard";
import { ReviewExpandedCard } from "../cards/ReviewExpandCard";

const ReviewsList: React.FC = () => {
  const { recentReviews, loading } = useRecentReviews();
  const [expandedReview, setExpandedReview] = useState<Review | null>(null);

  if (loading) return <p className="text-center text-gray-400">Загрузка...</p>;

  return (
    <section className="w-304 mx-auto py-16 px-4">
      <h2 className="text-3xl font-bold text-center mb-12">Свежие рецензии</h2>
      <div className="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-8">
        {recentReviews.map((review: Review) => (
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
    </section>
  );
};

export { ReviewsList };
