using Avalonia.Controls;
using System.Collections.Generic;
using System;
using System.IO;
using System.Threading.Tasks;
using System.Text;

namespace Blob_Editor.ViewModels
{
    public class MainWindowViewModel : ViewModelBase
    {
        private string filepath;
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

            this.filepath = result[0];
            Console.WriteLine($"Dialog directory: {this.filepath}");
            this.cfgViewModel.Elements.Clear();
            CFG cfg = new CFG(this.filepath);
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

        public void OnSaveClickCommand()
        {
            StreamWriter sw = new StreamWriter(this.filepath, false);
            foreach (Element element in this.cfgViewModel.Elements)
            {
                if (element is Entry)
                {
                    if (element.ValueList.Count == 0)
                    {
                        sw.WriteLine("{0} = ", element.Key);
                    } else
                    {
                        sw.WriteLine("{0} = {1}", element.Key, element.ValueList[0]);
                        for (int i = 1; i < element.ValueList.Count-1; i++)
                        {
                            sw.WriteLine("{0}", element.ValueList[i]);
                        }
                    }
                } else if (element is Comment) 
                {
                    sw.WriteLine(element.Key);
                } else if (element is Empty)
                {
                    sw.WriteLine();
                }
            }
            sw.Close();
        }

        public void OnAddClickCommand()
        {
            Console.WriteLine("Adding item...");
        }

        public void OnDeleteClickCommand()
        {
            Console.WriteLine("Deleting item...");
        }
    }
}
