﻿using Avalonia.Controls;
using System.Collections.Generic;
using System;
using System.Threading.Tasks;
using Blob_Editor;
namespace Blob_Editor.ViewModels
{
    public class MainWindowViewModel : ViewModelBase
    {
        public string Greeting => "Hello World!";
        public string Status => "Everything ok.";

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
            }
            else
            {
                Console.WriteLine($"Dialog directory: {result[0]}");
                KagConfig kagConfig = Parser.Parse(result[0]);
                foreach (Element element in kagConfig.Elements)
                {
                    Console.WriteLine(element.Print());
                }
            }
        }

        public void OnExitClickCommand()
        {
            Environment.Exit(0);
        }
    }
}
