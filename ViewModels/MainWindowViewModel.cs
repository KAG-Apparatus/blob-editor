using Avalonia.Controls;
using System.Collections.Generic;
using System;
using System.Threading.Tasks;

namespace Blob_Editor.ViewModels
{
    public class MainWindowViewModel : ViewModelBase
    {
        public string Greeting => "Hello World!";
        public string Status => "Everything ok.";
        public CFGViewModel cfgViewModel { get; }

        public MainWindowViewModel()
        {
            this.cfgViewModel = new CFGViewModel(new List<Element>());
        }

        public async void OnOpenClickCommand()
        {
            List<FileDialogFilter> filters = new List<FileDialogFilter>();
            filters.Add( //create filter for cfg files
                new FileDialogFilter()
                {
                    Extensions = new List<string>() { "cfg" },
                    Name = "KAG Config File"
                }
            );

            OpenFileDialog dialog = new OpenFileDialog()
            {
                AllowMultiple = false,
                Filters = filters,
                Title = "Open KAG Config File"
            };

            Task<string[]> task = dialog.ShowAsync(new Window());
            await task;
            string[] result = task.Result;


            if (result == null || result.Length != 1)
            {
                Console.WriteLine("Error opening file");
                return;
            }

            Console.WriteLine($"Dialog directory: {result[0]}");
            CFG cfg = new CFG(result[0]);
            foreach (Element element in cfg.Elements)
            {
                Console.WriteLine(element.Print());
                this.cfgViewModel.Elements.Add(element);
            }
        }

        public void OnExitClickCommand()
        {
            Environment.Exit(0);
        }

		public void OnClearClickCommand()
		{
			this.cfgViewModel.Elements.Clear();
		}
    }
}
