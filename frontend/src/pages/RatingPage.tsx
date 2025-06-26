import { useState } from "react";
import { Footer } from "../components/blocks/Footer";
import { Header } from "../components/blocks/Header";
import { Tabs } from "../components/ui/Tabs";
import { Spinner } from "../components/ui/Spinner";
import { TrackCard } from "../components/cards/TrackCard";
import ModelWindow from "../components/blocks/ModelWindow";
import type { Track } from "../types";
import { useTopTracks } from "../hooks/useTopTracks";
import { useReviewers } from "../hooks/useReviewers";
import { ReviewerCard } from "../components/cards/ReviewerCard";

type TabType = "week" | "month" | "all" | "reviewers";

const RatingPage = () => {
  const [activeTab, setActiveTab] = useState<TabType>("week");
  const [selectedTrack, setSelectedTrack] = useState<Track | null>(null);

  const { tracks, loading: tracksLoading } = useTopTracks(
    activeTab === "all" ? "all" : activeTab,
    activeTab === "all" ? 100 : 20,
  );

  const { reviewers, loading: reviewersLoading } = useReviewers();

  const tabs: Array<{ id: TabType; label: string }> = [
    { id: "week", label: "Топ недели" },
    { id: "month", label: "Топ месяца" },
    { id: "all", label: "Топ-100" },
    { id: "reviewers", label: "Лучшие рецензисты" },
  ];

  return (
    <>
      <Header />

      <main className="min-h-screen bg-gray-900 text-white pt-10 pb-5">
        <div className="w-304 mx-auto px-4">
          <div className="flex flex-col md:flex-row justify-between items-start md:items-center mb-8 gap-4">
            <h1 className="text-3xl font-bold">Рейтинги</h1>
          </div>

          <Tabs
            tabs={tabs}
            activeTab={activeTab}
            onChange={(tab: TabType) => {
              setActiveTab(tab);
              window.scrollTo(0, 0);
            }}
          />

          <div className="mt-8">
            {activeTab !== "reviewers" ? (
              tracksLoading ? (
                <div className="flex justify-center py-20">
                  <Spinner size="lg" />
                </div>
              ) : tracks.length === 0 ? (
                <div className="text-center py-10 text-gray-400">
                  Нет данных для отображения
                </div>
              ) : (
                <>
                  <div className="grid grid-cols-2 sm:grid-cols-3 md:grid-cols-4 lg:grid-cols-5 gap-6">
                    {tracks.map((track, index) => (
                      <div key={track.id} className="relative">
                        <div className="absolute -left-2 -top-2 bg-purple-600 text-white font-bold rounded-full w-6 h-6 flex items-center justify-center z-10">
                          {index + 1}
                        </div>
                        <TrackCard
                          track={track}
                          onClick={() => setSelectedTrack(track)}
                        />
                      </div>
                    ))}
                  </div>

                  {selectedTrack && (
                    <ModelWindow
                      track={selectedTrack}
                      onClose={() => setSelectedTrack(null)}
                    />
                  )}
                </>
              )
            ) : (
              <div className="mt-6">
                <h2 className="text-xl font-semibold mb-6">
                  Лучшие рецензисты
                </h2>
                {reviewersLoading ? (
                  <div className="flex justify-center py-20">
                    <Spinner size="lg" />
                  </div>
                ) : reviewers.length === 0 ? (
                  <div className="text-center py-10 text-gray-400">
                    Нет данных о рецензистах
                  </div>
                ) : (
                  <div className="grid grid-cols-1 sm:grid-cols-2 md:grid-cols-3 lg:grid-cols-4 gap-6">
                    {reviewers.map((reviewer, index) => (
                      <ReviewerCard
                        key={reviewer.id}
                        reviewer={reviewer}
                        rank={index + 1}
                      />
                    ))}
                  </div>
                )}
              </div>
            )}
          </div>
        </div>
      </main>

      <Footer />
    </>
  );
};

export default RatingPage;
