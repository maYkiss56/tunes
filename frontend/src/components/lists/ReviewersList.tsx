import { useReviewers } from "../../hooks/useReviewers";
import { ReviewerCard } from "../cards/ReviewerCard";
import { Spinner } from "../ui/Spinner";

export const ReviewersList = () => {
  const { reviewers, loading } = useReviewers();

  if (loading) {
    return (
      <div className="flex justify-center py-12">
        <Spinner size="lg" />
      </div>
    );
  }

  if (reviewers.length === 0) {
    return (
      <div className="rounded-lg bg-gray-100 p-6 text-center">
        <p className="text-gray-600">Пока нет рецензентов</p>
      </div>
    );
  }

  return (
    <div className="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-8">
      {reviewers.map((reviewer, index) => (
        <ReviewerCard key={reviewer.id} reviewer={reviewer} rank={index + 1} />
      ))}
    </div>
  );
};
