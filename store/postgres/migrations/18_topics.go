package migrations

func init() {
	include(18, query(`
		insert into topic (name, slug, created_at) values
			('Музыка', 'Music', now()),
			('Новости', 'News', now()),
			('СМИ', 'Mass media', now()),
			('Криптовалюты', 'Cryptocurrencies', now()),
			('Маркетинг', 'Marketing', now()),
			('Реклама', 'Ad', now()),
			('Видео', 'Video', now()),
			('Фильмы', 'Movies', now()),
			('Кино', 'Movie', now()),
			('Искусство', 'Art', now()),
			('Фото', 'Photo', now()),
			('Здоровье', 'Health', now()),
			('Спорт', 'Sport', now()),
			('Познавательное', 'Cognitive', now()),
			('Политика', 'politic', now()),
			('Технологии', 'Technologies', now()),
			('Образование', 'Education', now()),
			('Путешествия', 'Travels', now()),
			('Психология', 'Psychology', now()),
			('Цитаты', 'Quotes', now()),
			('Мода', 'Fashion', now()),
			('Красота', 'Beauty', now()),
			('Книги', 'Books', now()),
			('Игры', 'Games', now()),
			('Приложения', 'Apps', now()),
			('Экономика', 'Economy', now()),
			('Карьера', 'Career', now()),
			('Кулинария', 'Cooking', now()),
			('Еда', 'Food', now()),
			('Авто', 'Auto', now()),
			('Лингвистика', 'Linguistics', now()),
			('Дизайн', 'Design', now()),
			('Религия', 'Religion', now()),
			('Каталоги', 'Catalogs', now()),
			('Telegram', 'Telegram', now()),
			('Семья', 'Family', now()),
			('Дети', 'Kids', now()),
			('Животные', 'Animals', now()),
			('Медицина', 'Medicine', now()),
			('Рукоделие', 'Needlework', now()),
			('Лайфхаки', 'Lifehacks', now()),
			('ДляВзрослых', 'ForAdults', now()),
			('Другое', 'Other', now());
	`), query(`
	`))
}