import { Link } from "react-router-dom";
import { Footer } from "../components/blocks/Footer";
import { Header } from "../components/blocks/Header";
import { Button } from "../components/ui/Button";
import {
  CommunityIcon,
  DiscoveryIcon,
  MusicNoteIcon,
} from "../components/ui/icons";
import { FeatureCard } from "../components/cards/FeatureCard";
import { ReviewsList } from "../components/lists/ReviewsList";
import { ReviewersList } from "../components/lists/ReviewersList";

const Home = () => {
  return (
    <>
      <Header />
      <main>
        <section className="bg-gradient-to-r from-purple-900 to-black py-20 px-4 text-center">
          <h1 className="text-4xl md:text-6xl font-bold text-white mb-6">
            Открывайте. Слушайте. Оценивайте.
          </h1>
          <p className="text-xl text-gray-300 mb-8 max-w-3xl mx-auto">
            Присоединяйтесь к сообществу MelodyCritic и делитесь своими
            впечатлениями о последних музыкальных релизах
          </p>
          <div className="flex justify-center gap-4">
            <Link to="/login">
              <Button size="lg">Начать сейчас</Button>
            </Link>
            <Button variant="secondary" size="lg">
              Как это работает?
            </Button>
          </div>
        </section>

        <ReviewsList />

        <section className="bg-gray-100 py-16 px-4">
          <h2 className="text-3xl font-bold text-center mb-12 text-purple-600">
            Почему MelodyCritic?
          </h2>
          <div className="grid md:grid-cols-3 gap-8 max-w-6xl mx-auto">
            <FeatureCard
              icon={<MusicNoteIcon />}
              title="Экспертные оценки"
              text="Получайте профессиональные рецензии от музыкальных критиков"
            />
            <FeatureCard
              icon={<CommunityIcon />}
              title="Живое сообщество"
              text="Обсуждайте музыку с единомышленниками"
            />
            <FeatureCard
              icon={<DiscoveryIcon />}
              title="Персональные рекомендации"
              text="Открывайте новую музыку по вашим вкусам"
            />
          </div>
        </section>

        <section className="w-304 mx-auto py-16 px-4">
          <h2 className="text-3xl font-bold text-center mb-12">
            Наши топ-рецензенты
          </h2>
          <ReviewersList />
        </section>

        <section className="bg-gradient-to-br from-purple-900 to-pink-800 py-20 px-4 text-center">
          <h2 className="text-3xl font-bold text-white mb-6">
            Присоединяйтесь к MelodyCritic сегодня
          </h2>
          <p className="text-xl text-purple-100 mb-8 max-w-2xl mx-auto">
            Зарегистрируйтесь, чтобы сохранять любимые треки, писать рецензии и
            получать персональные рекомендации
          </p>
          <Link to="/register">
            <Button variant="primary" size="lg" className="hover:bg-gray-100">
              Создать аккаунт
            </Button>
          </Link>
        </section>
      </main>
      <Footer />
    </>
  );
};

export default Home;
