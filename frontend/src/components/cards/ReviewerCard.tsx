import type { Reviewer } from "../../types";

interface ReviewerCardProps {
  reviewer: Reviewer;
  rank?: number;
}

export const ReviewerCard = ({ reviewer, rank }: ReviewerCardProps) => {
  return (
    <div className="bg-white rounded-lg overflow-hidden shadow-md hover:shadow-xl transition-all duration-300 border border-gray-100">
      <div className="relative bg-gradient-to-r from-purple-50 to-blue-50 p-4">
        {rank && (
          <div className="absolute top-3 left-3 bg-white rounded-full w-8 h-8 flex items-center justify-center shadow-sm">
            <span className="font-bold text-purple-600">{rank}</span>
          </div>
        )}

        <div className="flex justify-center mt-2">
          {reviewer.avatar_url ? (
            <img
              src={`http://localhost:8080/${reviewer.avatar_url}`}
              alt={reviewer.username}
              className="w-20 h-20 rounded-full object-cover border-4 border-white shadow-md"
            />
          ) : (
            <div className="w-20 h-20 rounded-full bg-gradient-to-br from-purple-400 to-blue-500 flex items-center justify-center text-white text-3xl font-bold shadow-md">
              {reviewer.username.charAt(0).toUpperCase()}
            </div>
          )}
        </div>
      </div>

      <div className="p-4 text-center">
        <h3 className="text-lg font-bold text-gray-800 mb-1">
          {reviewer.username}
        </h3>

        <div className="flex items-center justify-center space-x-2 mb-3">
          <span className="bg-purple-100 text-purple-800 text-xs font-semibold px-2 py-1 rounded-full">
            {reviewer.review_count}
          </span>

          <span className="text-sm text-gray-500">
            {getReviewWord(reviewer.review_count)}
          </span>
        </div>

        <div className="flex justify-center">
          <div className="inline-flex items-center bg-blue-50 text-blue-700 px-3 py-1 rounded-full text-xs font-medium">
            <svg
              className="w-3 h-3 mr-1"
              fill="currentColor"
              viewBox="0 0 20 20"
            >
              <path d="M9.049 2.927c.3-.921 1.603-.921 1.902 0l1.07 3.292a1 1 0 00.95.69h3.462c.969 0 1.371 1.24.588 1.81l-2.8 2.034a1 1 0 00-.364 1.118l1.07 3.292c.3.921-.755 1.688-1.54 1.118l-2.8-2.034a1 1 0 00-1.175 0l-2.8 2.034c-.784.57-1.838-.197-1.539-1.118l1.07-3.292a1 1 0 00-.364-1.118L2.98 8.72c-.783-.57-.38-1.81.588-1.81h3.461a1 1 0 00.951-.69l1.07-3.292z" />
            </svg>
            Топ рецензент
          </div>
        </div>
      </div>
    </div>
  );
};

function getReviewWord(count: number): string {
  const lastDigit = count % 10;
  const lastTwoDigits = count % 100;

  if (lastTwoDigits >= 11 && lastTwoDigits <= 19) return "рецензий";
  if (lastDigit === 1) return "рецензия";
  if (lastDigit >= 2 && lastDigit <= 4) return "рецензии";
  return "рецензий";
}
