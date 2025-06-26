import { type FC, type ReactNode } from "react";

type FeatureCardProps = {
  icon: ReactNode;
  title: string;
  text: string;
};

const FeatureCard: FC<FeatureCardProps> = ({ icon, title, text }) => {
  return (
    <div className="bg-white p-6 rounded-lg shadow-md hover:shadow-lg transition-shadow duration-300 h-full">
      <div className="text-purple-500 text-4xl mb-4">{icon}</div>
      <h3 className="text-xl font-bold text-gray-900 mb-2">{title}</h3>
      <p className="text-gray-600">{text}</p>
    </div>
  );
};

export { FeatureCard };
