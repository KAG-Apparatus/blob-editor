using Avalonia;
using Avalonia.Controls;
using Avalonia.Markup.Xaml;

namespace Blob_Editor.Views
{
    public class CFGView : UserControl
    {
        public CFGView()
        {
            InitializeComponent();
        }

        private void InitializeComponent()
        {
            AvaloniaXamlLoader.Load(this);
        }
    }
}