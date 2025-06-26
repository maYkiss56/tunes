import type { FC } from "react";
import { Header } from "../components/blocks/Header";
import { Footer } from "../components/blocks/Footer";
import { ProfileCard } from "../components/cards/ProfileCard";
import { ReviewsStatsCard } from "../components/cards/ReviewsStatCard";

const ProfilePage: FC = () => {
  return (
    <>
      <Header />
      <div className=" bg-gray-900">
        <div className="w-304 mx-auto px-4 py-8">
          <div className="flex justify-between items-center">
            <ProfileCard />
            <ReviewsStatsCard />
          </div>
        </div>
      </div>
      <Footer />
    </>
  );
};

export default ProfilePage;
