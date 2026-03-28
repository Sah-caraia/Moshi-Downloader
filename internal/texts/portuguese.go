package texts

import "moshi/internal/models"

var PortugueseErrors = models.ErrText{
	NotFoundLibrary: "O programa \"yt-dlp\" não está instalado! Para instalá-lo, primeiro instale o <Python>. Logo após, use: \"pip install -U yt-dlp\" e rode o programa novamente.",
	NetworkError: "Parece que você não está conectado à internet :/\nTente novamente mais tarde.",
	AccessError: "Não foi possível salvar o arquivo devido um erro de permissão.",
	DownloadError: "Ocorreu um erro ao baixar o vídeo! Talvez a mídia não esteja disponível.",
	InvalidUrl: "Envie uma URL válida!",
}

var PortugueseSuccess = models.SusText{
	MediaInstalled: "Mídia instalada ✓",
	AllSuccess: "Todas as mídias foram instaladas com sucesso! ✓",
}

var PortugueseDefault = models.DefText{
	UseKeyboardArrow: "Use as setas do teclado (↑↓←→) para navegar entre as opções",
	SelectAnOption: "Selecione uma das opções abaixo: ",
	SendYourUrl: "Envie a sua URL",
	InputHint: "Quando acabar de enviar as URLs, digite \"END\" e pressione <ENTER>.",
}
