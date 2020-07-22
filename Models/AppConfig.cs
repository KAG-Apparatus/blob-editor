using System.IO;
using Newtonsoft.Json;
using Avalonia.Controls;
using System.Collections.Generic;
namespace Blob_Editor
{
	struct AppConfig
	{
		public Window window;
		public static AppConfig GenerateFromRaw(string raw, Window window)
		{
			AppConfig config = JsonConvert.DeserializeObject<AppConfig>(raw);
			config.window = window;
			return config;
		}
		public static AppConfig GenerateFromFilePath(string filePath, Window window)
		{
			string raw = File.ReadAllText(filePath);

			return GenerateFromRaw(raw, window);
		}
		public int resX,resY;
		public string defaultOpenDir; 

		public void AddControls()
		{
			TabItem item = window.FindControl<TabItem>("Settings_TabItem"); //not sure what to do from here to add items;
		}
	}
}